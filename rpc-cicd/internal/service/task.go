package service

import (
	"context"

	"github.com/hatlonely/go-rpc/rpc-cicd/api/gen/go/api"
)

func (s *CICDService) GetTask(ctx context.Context, req *api.GetTaskReq) (*api.Task, error) {
	return s.storage.GetTask(ctx, req.Id)
}

func (s *CICDService) DelTask(ctx context.Context, req *api.DelTaskReq) (*api.Empty, error) {
	return &api.Empty{}, s.storage.DelTask(ctx, req.Id)
}

func (s *CICDService) PutTask(ctx context.Context, req *api.PutTaskReq) (*api.Empty, error) {
	return &api.Empty{}, s.storage.PutTask(ctx, req.Task)
}

func (s *CICDService) UpdateTask(ctx context.Context, req *api.UpdateTaskReq) (*api.Empty, error) {
	return &api.Empty{}, s.storage.UpdateTask(ctx, req.Task)
}

func (s *CICDService) ListTask(ctx context.Context, req *api.ListTaskReq) (*api.ListTaskRes, error) {
	res, err := s.storage.ListTask(ctx, req.Offset, req.Limit)
	if err != nil {
		return nil, err
	}
	return &api.ListTaskRes{Tasks: res}, nil
}
