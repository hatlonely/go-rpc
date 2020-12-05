package service

import (
	"bytes"
	"context"
	"encoding/json"
	"os/exec"
	"syscall"
	"text/template"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/hatlonely/go-rpc/rpc-cicd/api/gen/go/api"
)

const JobStatusRunning = "Running"
const JobStatusWaiting = "Waiting"
const JobStatusFailed = "Failed"
const JobStatusFinish = "Finish"

func (s *CICDService) RunTask(ctx context.Context, req *api.RunTaskReq) (*api.RunTaskRes, error) {
	// 获取 task
	task, err := s.storage.GetTask(ctx, req.TaskID)
	if err != nil {
		return nil, err
	}

	jobID := primitive.NewObjectID()
	job := api.Job{
		Id:     jobID.String(),
		TaskID: req.TaskID,
		Status: JobStatusRunning,
	}

	// 插入 job
	if err := s.storage.PutJob(ctx, &job); err != nil {
		return nil, err
	}

	// 执行 job
	go func() {
		_ = s.runTask(ctx, jobID.String(), task)
	}()

	return &api.RunTaskRes{JobID: jobID.String()}, nil
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

func (s *CICDService) runTask(ctx context.Context, jobID string, task *api.Task) error {
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

	job, err := s.storage.GetJob(ctx, jobID)
	if err != nil {
		return err
	}
	for _, i := range templates {
		tpl, err := template.New("").Parse(i.ScriptTemplate.Script)
		if err != nil {
			return errors.Wrap(err, "create template failed")
		}

		buf := &bytes.Buffer{}
		if err := tpl.Execute(buf, kvs); err != nil {
			return errors.Wrap(err, "tpl execute failed")
		}

		exitCode, stdout, stderr, err := Exec(i.ScriptTemplate.Language, buf.String())
		if err != nil {
			return errors.Wrap(err, "Exec failed")
		}

		job.Subs = append(job.Subs, &api.Job_Sub{
			TemplateID: i.Id,
			ExitCode:   int32(exitCode),
			Stdout:     stdout,
			Stderr:     stderr,
			Status:     "Success",
		})

		if err := s.storage.UpdateJob(ctx, job); err != nil {
			return err
		}
	}

	job.Status = JobStatusFailed
	if err := s.storage.UpdateJob(ctx, job); err != nil {
		return err
	}

	return nil
}

func Exec(interpreter string, script string) (int, string, string, error) {
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
