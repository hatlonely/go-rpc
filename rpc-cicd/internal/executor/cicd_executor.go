package executor

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"sync"
	"syscall"
	"time"

	"github.com/cbroglie/mustache"
	"github.com/hatlonely/go-kit/logger"
	"github.com/pkg/errors"

	"github.com/hatlonely/go-rpc/rpc-cicd/api/gen/go/api"
	"github.com/hatlonely/go-rpc/rpc-cicd/internal/storage"
)

type CICDExecutor struct {
	storage *storage.CICDStorage
	options *CICDOptions

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup

	log *logger.Logger
}

func NewCICDExecutorWithOptions(storage *storage.CICDStorage, options *CICDOptions) *CICDExecutor {
	ctx, cancel := context.WithCancel(context.Background())

	return &CICDExecutor{
		storage: storage,
		options: options,
		ctx:     ctx,
		cancel:  cancel,
		log:     logger.NewStdoutJsonLogger(),
	}
}

func (e *CICDExecutor) SetLogger(log *logger.Logger) {
	e.log = log
}

func (e *CICDExecutor) Run() {
	for i := 0; i < e.options.WorkerNum; i++ {
		e.wg.Add(1)
		go func() {
			e.runLoop(e.ctx)
			e.wg.Done()
		}()
	}
}

func (e *CICDExecutor) Stop() {
	e.cancel()
	e.wg.Wait()
}

type CICDOptions struct {
	WorkerNum int           `dft:"20"`
	Data      string        `dft:"data"`
	SleepTime time.Duration `dft:"5s"`
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

func (e *CICDExecutor) runLoop(ctx context.Context) {
	errs := make(chan error, 1)

out:
	for {
		go func() {
			errs <- e.runOnce(NewExecutorContext(ctx))
		}()
		select {
		case <-ctx.Done():
			break out
		case <-errs:
		}
	}
}

func (e *CICDExecutor) runOnce(ctx context.Context) error {
	now := time.Now()
	var err error
	var job *api.Job
	var acquire bool
	defer func() {
		if acquire {
			e.log.Info(map[string]interface{}{
				"id":       job.Id,
				"err":      err,
				"errStack": fmt.Sprintf("%+v", err),
				"timeMs":   time.Now().Sub(now).Milliseconds(),
				"ctx":      ctx.Value(executorKey{}),
			})
		}
	}()

	job, err = e.storage.FindOneUnscheduleJob(ctx)
	if err != nil {
		return err
	}
	if job == nil {
		select {
		case <-ctx.Done():
		case <-time.After(e.options.SleepTime):
		}
		return nil
	}
	acquire, err = e.storage.UpdateJobStatus(ctx, job.Id, storage.JobStatusWaiting, storage.JobStatusWaiting)
	if err != nil {
		return err
	}
	if !acquire {
		return nil
	}
	err = e.runTask(ctx, job.Id)
	return err
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

	job.Status = storage.JobStatusRunning
	job.ScheduleAt = int32(time.Now().Unix())
	if err := e.storage.UpdateJob(ctx, job); err != nil {
		return err
	}

	now := time.Now()
	if err := e.runSubTasks(ctx, job, task); err != nil {
		job.Error = err.Error()
		job.Status = storage.JobStatusFailed
	} else {
		job.Status = storage.JobStatusFinish
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
			Status:       storage.JobStatusWaiting,
			UpdateAt:     int32(time.Now().Unix()),
		})
	}

	if err := e.storage.UpdateJob(ctx, job); err != nil {
		return err
	}

	for _, sub := range job.Subs {
		if sub.Status == storage.JobStatusFinish {
			continue
		}

		sub.Status = storage.JobStatusRunning
		if err := e.storage.UpdateJob(ctx, job); err != nil {
			return err
		}

		now := time.Now()
		exitCode, stdout, stderr, err := Exec(sub.Language, sub.Script, fmt.Sprintf("%v/%v/%v", e.options.Data, job.TaskName, job.Seq))
		if err != nil {
			return errors.Wrapf(err, "exec [%v] failed", sub.TemplateName)
		}

		sub.Status = storage.JobStatusFailed
		if exitCode == 0 {
			sub.Status = storage.JobStatusFinish
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
