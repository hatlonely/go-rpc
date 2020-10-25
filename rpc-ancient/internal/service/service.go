package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/gorm"

	"github.com/hatlonely/go-rpc/rpc-ancient/api/gen/go/api"
)

type AncientService struct {
	mysqlCli *gorm.DB
}

func NewAncientService(mysqlCli *gorm.DB) (*AncientService, error) {
	return &AncientService{
		mysqlCli: mysqlCli,
	}, nil
}

func (s *AncientService) GetAncient(ctx context.Context, req *api.GetAncientReq) (*api.Ancient, error) {
	return nil, nil
}

func (s *AncientService) PutAncient(ctx context.Context, req *api.PutAncientReq) (*empty.Empty, error) {
	return nil, nil
}

func (s *AncientService) UpdateAncient(ctx context.Context, in *api.UpdateAncientReq) (*empty.Empty, error) {
	return nil, nil
}
