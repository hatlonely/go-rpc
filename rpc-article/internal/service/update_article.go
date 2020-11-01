package service

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hatlonely/go-kit/rpcx"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/hatlonely/go-rpc/rpc-article/api/gen/go/api"
	"github.com/hatlonely/go-rpc/rpc-article/internal/storage"
)

func (s *ArticleService) UpdateArticle(ctx context.Context, req *api.UpdateArticleReq) (*empty.Empty, error) {
	article := &storage.Article{
		ID:      int(req.Article.Id),
		OwnerID: int(req.Article.OwnerID),
		Title:   req.Article.Title,
		Tags:    req.Article.Tags,
		Brief:   req.Article.Tags,
		Content: req.Article.Content,
		UTime:   time.Now(),
	}

	rpcx.CtxSet(ctx, "article", article)

	if err := s.mysqlCli.Update(article).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			if err := s.mysqlCli.Create(article).Error; err != nil {
				return nil, errors.Wrap(err, "mysql update failed")
			}
		} else {
			return nil, errors.Wrap(err, "mysql update failed")
		}
	}

	return &empty.Empty{}, nil
}
