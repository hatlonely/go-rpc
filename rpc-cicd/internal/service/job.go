package service

import (
	"context"

	"github.com/hatlonely/go-rpc/rpc-cicd/api/gen/go/api"
)

func (s *CICDService) GetJob(ctx context.Context, req *api.GetJobReq) (*api.Job, error) {
	return s.storage.GetJob(ctx, req.Id)
}

func (s *CICDService) DelJob(ctx context.Context, req *api.DelJobReq) (*api.DelJobRes, error) {
	return &api.DelJobRes{Id: req.Id}, s.storage.DelJob(ctx, req.Id)
}

func (s *CICDService) ListJob(ctx context.Context, req *api.ListJobReq) (*api.ListJobRes, error) {
	res, err := s.storage.ListJobWithTaskID(ctx, req.TaskID, req.Offset, req.Limit)
	if err != nil {
		return nil, err
	}
	return &api.ListJobRes{Jobs: res}, nil
}
