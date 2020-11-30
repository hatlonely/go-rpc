package service

import (
	"context"

	"github.com/hatlonely/go-rpc/rpc-cicd/api/gen/go/api"
)

func (s *CICDService) RunTask(ctx context.Context, req *api.RunTaskReq) (*api.RunTaskRes, error) {

	return &api.RunTaskRes{}, nil
}
