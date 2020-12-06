package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"text/template"
	"time"

	"github.com/hatlonely/go-kit/rpcx"
	"github.com/pkg/errors"

	"github.com/hatlonely/go-rpc/rpc-cicd/api/gen/go/api"
	"github.com/hatlonely/go-rpc/rpc-cicd/internal/executor"
)

const JobStatusRunning = "Running"
const JobStatusWaiting = "Waiting"
const JobStatusFailed = "Failed"
const JobStatusFinish = "Finish"

func (s *CICDService) RunTask(ctx context.Context, req *api.RunTaskReq) (*api.RunTaskRes, error) {
	job := api.Job{
		TaskID:   req.TaskID,
		Status:   JobStatusWaiting,
		CreateAt: int32(time.Now().Unix()),
	}

	jobID, err := s.storage.PutJob(ctx, &job)
	if err != nil {
		return nil, err
	}

	s.executor.AddTask(rpcx.MetaDataGetRequestID(ctx), jobID)

	return &api.RunTaskRes{JobID: jobID}, nil
}

func (s *CICDService) GetTemplates(ctx context.Context, req *api.GetTemplatesReq) (*api.ListTemplateRes, error) {
	res, err := s.storage.GetTemplateByIDs(ctx, req.Ids)
	if err != nil {
		return nil, err
	}

	return &api.ListTemplateRes{Templates: res}, nil
}

func (s *CICDService) GetVariables(ctx context.Context, req *api.GetVariablesReq) (*api.ListVariableRes, error) {
	res, err := s.storage.GetVariableByIDs(ctx, req.Ids)
	if err != nil {
		return nil, err
	}

	return &api.ListVariableRes{Variables: res}, nil
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

func (s *CICDService) ExecutorHandler(ctx context.Context, jobID interface{}) error {
	return s.runTask(ctx, jobID.(string))
}

func (s *CICDService) runTask(ctx context.Context, jobID string) error {
	job, err := s.storage.GetJob(ctx, jobID)
	if err != nil {
		return err
	}
	defer executor.CtxSet(ctx, "job", job)

	task, err := s.storage.GetTask(ctx, job.TaskID)
	if err != nil {
		return err
	}

	job.TaskName = task.Name
	job.Status = JobStatusRunning
	if err := s.storage.UpdateJob(ctx, job); err != nil {
		return err
	}
	defer executor.CtxSet(ctx, "task", task)

	if err := s.runSubTasks(ctx, job, task); err != nil {
		job.Error = err.Error()
		job.Status = JobStatusFailed
	} else {
		job.Status = JobStatusFinish
	}
	if err := s.storage.UpdateJob(ctx, job); err != nil {
		return err
	}
	return nil
}

func (s *CICDService) runSubTasks(ctx context.Context, job *api.Job, task *api.Task) error {
	variables, err := s.storage.GetVariableByIDs(ctx, task.VariableIDs)
	if err != nil {
		return errors.Wrap(err, "GetVariables failed")
	}
	templates, err := s.storage.GetTemplateByIDs(ctx, task.TemplateIDs)
	if err != nil {
		return errors.Wrap(err, "GetTemplates failed")
	}
	kvs, err := mergeVariables(variables)
	if err != nil {
		return errors.Wrap(err, "mergeVariables failed")
	}

	for _, i := range templates {
		tpl, err := template.New("").Parse(i.ScriptTemplate.Script)
		if err != nil {
			return errors.Wrapf(err, "create template [%v] failed", i.Name)
		}

		buf := &bytes.Buffer{}
		if err := tpl.Execute(buf, kvs); err != nil {
			return errors.Wrapf(err, "tpl execute [%v] failed", i.Name)
		}

		exitCode, stdout, stderr, err := Exec(i.ScriptTemplate.Language, buf.String(), fmt.Sprintf("%v/%v", s.options.Data, job.Id))
		if err != nil {
			return errors.Wrapf(err, "Exec [%v] failed", i.Name)
		}

		job.Subs = append(job.Subs, &api.Job_Sub{
			TemplateID:   i.Id,
			TemplateName: i.Name,
			ExitCode:     int32(exitCode),
			Language:     i.ScriptTemplate.Language,
			Script:       buf.String(),
			Stdout:       stdout,
			Stderr:       stderr,
			Status:       "Success",
		})

		if err := s.storage.UpdateJob(ctx, job); err != nil {
			return err
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
