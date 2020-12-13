package service

import (
	"context"
	"time"

	"github.com/hatlonely/go-rpc/rpc-cicd/api/gen/go/api"
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

	//s.executor.AddTask(rpcx.MetaDataGetRequestID(ctx), jobID)

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
