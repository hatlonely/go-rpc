package service

import (
	"context"

	"github.com/hatlonely/go-rpc/rpc-cicd/api/gen/go/api"
)

func (s *CICDService) GetVariable(ctx context.Context, req *api.GetVariableReq) (*api.Variable, error) {
	return s.storage.GetVariable(ctx, req.Id)
}

func (s *CICDService) DelVariable(ctx context.Context, req *api.DelVariableReq) (*api.DelVariableRes, error) {
	return &api.DelVariableRes{Id: req.Id}, s.storage.DelVariable(ctx, req.Id)
}

func (s *CICDService) PutVariable(ctx context.Context, req *api.PutVariableReq) (*api.PutVariableRes, error) {
	id, err := s.storage.PutVariable(ctx, req.Variable)
	if err != nil {
		return nil, err
	}
	return &api.PutVariableRes{Id: id}, nil
}

func (s *CICDService) UpdateVariable(ctx context.Context, req *api.UpdateVariableReq) (*api.UpdateVariableRes, error) {
	return &api.UpdateVariableRes{Id: req.Variable.Id}, s.storage.UpdateVariable(ctx, req.Variable)
}

func (s *CICDService) ListVariable(ctx context.Context, req *api.ListVariableReq) (*api.ListVariableRes, error) {
	res, err := s.storage.ListVariable(ctx, req.Offset, req.Limit)
	if err != nil {
		return nil, err
	}
	return &api.ListVariableRes{Variables: res}, nil
}
