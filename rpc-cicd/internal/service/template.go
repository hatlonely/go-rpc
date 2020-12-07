package service

import (
	"context"

	"github.com/hatlonely/go-rpc/rpc-cicd/api/gen/go/api"
)

func (s *CICDService) GetTemplate(ctx context.Context, req *api.GetTemplateReq) (*api.Template, error) {
	return s.storage.GetTemplate(ctx, req.Id)
}

func (s *CICDService) DelTemplate(ctx context.Context, req *api.DelTemplateReq) (*api.DelTemplateRes, error) {
	return &api.DelTemplateRes{Id: req.Id}, s.storage.DelTemplate(ctx, req.Id)
}

func (s *CICDService) PutTemplate(ctx context.Context, req *api.PutTemplateReq) (*api.PutTemplateRes, error) {
	id, err := s.storage.PutTemplate(ctx, req.Template)
	if err != nil {
		return nil, err
	}
	return &api.PutTemplateRes{Id: id}, nil
}

func (s *CICDService) UpdateTemplate(ctx context.Context, req *api.UpdateTemplateReq) (*api.UpdateTemplateRes, error) {
	return &api.UpdateTemplateRes{Id: req.Template.Id}, s.storage.UpdateTemplate(ctx, req.Template)
}

func (s *CICDService) ListTemplate(ctx context.Context, req *api.ListTemplateReq) (*api.ListTemplateRes, error) {
	res, err := s.storage.ListTemplate(ctx, req.Offset, req.Limit)
	if err != nil {
		return nil, err
	}
	return &api.ListTemplateRes{Templates: res}, nil
}
