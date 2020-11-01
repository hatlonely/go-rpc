package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"

	"github.com/hatlonely/go-rpc/rpc-article/api/gen/go/api"
	"github.com/hatlonely/go-rpc/rpc-article/internal/storage"
)

func (s *ArticleService) DelArticle(ctx context.Context, req *api.DelArticleReq) (*empty.Empty, error) {
	if err := s.mysqlCli.Delete(&storage.Article{}, req.Id).Error; err != nil {
		return nil, errors.Wrapf(err, "mysql delete article [%v] failed", req.Id)
	}

	return &empty.Empty{}, nil
}
