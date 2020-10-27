package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hatlonely/go-kit/rpcx"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/hatlonely/go-rpc/rpc-ancient/api/gen/go/api"
	"github.com/hatlonely/go-rpc/rpc-ancient/internal/model"
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
	requestID := rpcx.MetaDataGetRequestID(ctx)

	shici := &model.ShiCi{}
	if err := s.mysqlCli.Where("id=?", req.Id).First(shici).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, rpcx.NewErrorWithoutReferf(codes.NotFound, requestID, "NotFound", "shici [%v] not exist", req.Id)
		}
		return nil, errors.Wrapf(err, "mysql select shici [%v] failed", req.Id)
	}

	rpcx.CtxSet(ctx, "shici", shici)

	return &api.Ancient{
		Id:      int64(shici.ID),
		Title:   shici.Title,
		Author:  shici.Author,
		Dynasty: shici.Dynasty,
		Content: shici.Content,
	}, nil
}

func (s *AncientService) PutAncient(ctx context.Context, req *api.PutAncientReq) (*empty.Empty, error) {
	shici := &model.ShiCi{
		ID:      int(req.Ancient.Id),
		Title:   req.Ancient.Title,
		Author:  req.Ancient.Author,
		Dynasty: req.Ancient.Dynasty,
		Content: req.Ancient.Content,
	}

	rpcx.CtxSet(ctx, "shici", shici)

	if err := s.mysqlCli.Create(shici).Error; err != nil {
		return nil, errors.Wrap(err, "mysql create failed")
	}

	return nil, nil
}

func (s *AncientService) UpdateAncient(ctx context.Context, req *api.UpdateAncientReq) (*empty.Empty, error) {
	shici := &model.ShiCi{
		ID:      int(req.Ancient.Id),
		Title:   req.Ancient.Title,
		Author:  req.Ancient.Author,
		Dynasty: req.Ancient.Dynasty,
		Content: req.Ancient.Content,
	}

	rpcx.CtxSet(ctx, "shici", shici)

	if err := s.mysqlCli.Update(shici).Error; err != nil {
		return nil, errors.Wrap(err, "mysql update failed")
	}

	return nil, nil
}
