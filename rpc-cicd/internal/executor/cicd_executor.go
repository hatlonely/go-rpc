package executor

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/cbroglie/mustache"
	"github.com/hatlonely/go-kit/logger"
	"github.com/pkg/errors"

	"github.com/hatlonely/go-rpc/rpc-cicd/api/gen/go/api"
	"github.com/hatlonely/go-rpc/rpc-cicd/internal/storage"
)

const JobStatusRunning = "Running"
const JobStatusWaiting = "Waiting"
const JobStatusFailed = "Failed"
const JobStatusFinish = "Finish"

type CICDExecutor struct {
	executor *Executor
	storage  *storage.CICDStorage
	options  *CICDOptions
}

func NewCICDExecutorWithOptions(storage *storage.CICDStorage, options *CICDOptions) *CICDExecutor {
	ce := &CICDExecutor{
		storage: storage,
		options: options,
	}
	ce.executor = NewExecutorWithOptions(ce.ExecutorHandler, &options.Executor)

	return ce
}

func (e *CICDExecutor) SetLogger(log *logger.Logger) {
	e.executor.logger = log
}

func (e *CICDExecutor) Run() {
	e.executor.Run()
}

func (e *CICDExecutor) Stop() {
	e.executor.Stop()
}

type CICDOptions struct {
	Executor Options
	Data     string        `dft:"data"`
	Interval time.Duration `dft:"5s"`
}

func (e *CICDExecutor) listWaitingTask(ctx context.Context) {
	ticker := time.Tick(e.options.Interval)

out:
	for {
		select {
		case <-ctx.Done():
			break out
		case <-ticker:
			//jobs, e.storage.ListJob(ctx, "", 0, 20)
		}
	}
}

func mergeVariables(variables []*api.Variable) (map[string]interface{}, error) {
	kvs := map[string]interface{}{}

	for _, variable := range variables {
		var v interface{}
		if err := json.Unmarshal([]byte(variable.Kvs), &v); err != nil {
			return nil, errors.Wrap(err, "mergeVariables failed")
		}

		kvs[variable.Name] = v
	}

	return kvs, nil
}

func (e *CICDExecutor) ExecutorHandler(ctx context.Context, jobID interface{}) error {
	return e.runTask(ctx, jobID.(string))
}

func (e *CICDExecutor) runTask(ctx context.Context, jobID string) error {
	job, err := e.storage.GetJob(ctx, jobID)
	if err != nil {
		return err
	}
	defer CtxSet(ctx, "job", job)

	task, err := e.storage.GetTask(ctx, job.TaskID)
	if err != nil {
		return err
	}
	job.TaskName = task.Name
	defer CtxSet(ctx, "task", task)

	job.Status = JobStatusRunning
	job.ScheduleAt = int32(time.Now().Unix())
	if err := e.storage.UpdateJob(ctx, job); err != nil {
		return err
	}

	now := time.Now()
	if err := e.runSubTasks(ctx, job, task); err != nil {
		job.Error = err.Error()
		job.Status = JobStatusFailed
	} else {
		job.Status = JobStatusFinish
	}
	job.ElapseSeconds = int32(time.Now().Sub(now).Seconds())
	if err := e.storage.UpdateJob(ctx, job); err != nil {
		return err
	}
	return nil
}

func (e *CICDExecutor) runSubTasks(ctx context.Context, job *api.Job, task *api.Task) error {
	variables, err := e.storage.GetVariableByIDs(ctx, task.VariableIDs)
	if err != nil {
		return errors.Wrap(err, "GetVariables failed")
	}
	templates, err := e.storage.GetTemplateByIDs(ctx, task.TemplateIDs)
	if err != nil {
		return errors.Wrap(err, "GetTemplates failed")
	}
	m := map[string]*api.Template{}
	for _, t := range templates {
		m[t.Id] = t
	}
	for i, tid := range task.TemplateIDs {
		templates[i] = m[tid]
	}

	kvs, err := mergeVariables(variables)
	if err != nil {
		return errors.Wrap(err, "mergeVariables failed")
	}

	for _, i := range templates {
		str, err := mustache.Render(i.ScriptTemplate.Script, kvs)
		if err != nil {
			return errors.Wrapf(err, "mustache.Render failed. template: [%v]", i.Name)
		}

		job.Subs = append(job.Subs, &api.Job_Sub{
			TemplateID:   i.Id,
			TemplateName: i.Name,
			Language:     i.ScriptTemplate.Language,
			Script:       str,
			Status:       JobStatusWaiting,
			UpdateAt:     int32(time.Now().Unix()),
		})
	}

	if err := e.storage.UpdateJob(ctx, job); err != nil {
		return err
	}

	for _, sub := range job.Subs {
		sub.Status = JobStatusRunning
		if err := e.storage.UpdateJob(ctx, job); err != nil {
			return err
		}

		now := time.Now()
		exitCode, stdout, stderr, err := Exec(sub.Language, sub.Script, fmt.Sprintf("%v/%v/%v", e.options.Data, job.TaskName, job.Seq))
		if err != nil {
			return errors.Wrapf(err, "exec [%v] failed", sub.TemplateName)
		}

		sub.Status = JobStatusFailed
		if exitCode == 0 {
			sub.Status = JobStatusFinish
		}
		sub.Stdout = stdout
		sub.Stderr = stderr
		sub.ExitCode = int32(exitCode)
		sub.UpdateAt = int32(time.Now().Unix())
		sub.ElapseSeconds = int32(time.Now().Sub(now).Seconds())

		if err := e.storage.UpdateJob(ctx, job); err != nil {
			return err
		}

		if exitCode != 0 {
			return errors.Errorf("exit code %v", sub.ExitCode)
		}
	}

	return nil
}

func Exec(interpreter string, script string, workdir string) (int, string, string, error) {
	if err := os.MkdirAll(workdir, 0755); err != nil {
		return 0, "", "", errors.Wrap(err, "os.MkdirAll failed")
	}

	var stdout = &bytes.Buffer{}
	var stderr = &bytes.Buffer{}
	var cmd *exec.Cmd
	switch interpreter {
	case "python3":
		cmd = exec.Command("python3", "-c", script)
	case "bash", "shell", "sh":
		cmd = exec.Command("bash", "-c", script)
	default:
		return -1, "", "", errors.New("unsupported interpreter")
	}
	cmd.Dir = workdir
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	if err := cmd.Start(); err != nil {
		return -1, "", "", err
	}

	if err := cmd.Wait(); err != nil {
		if e, ok := err.(*exec.ExitError); ok {
			if status, ok := e.Sys().(syscall.WaitStatus); ok {
				return status.ExitStatus(), stdout.String(), stderr.String(), nil
			}
		}

		return -1, stdout.String(), stderr.String(), err
	}

	return 0, stdout.String(), stderr.String(), nil
}
