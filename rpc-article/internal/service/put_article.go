package service

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hatlonely/go-kit/rpcx"
	"github.com/pkg/errors"

	"github.com/hatlonely/go-rpc/rpc-article/api/gen/go/api"
	"github.com/hatlonely/go-rpc/rpc-article/internal/storage"
)

func (s *ArticleService) PutArticle(ctx context.Context, req *api.PutArticleReq) (*empty.Empty, error) {
	article := &storage.Article{
		ID:      int(req.Article.Id),
		OwnerID: int(req.Article.OwnerID),
		Title:   req.Article.Title,
		Tags:    req.Article.Tags,
		Brief:   req.Article.Tags,
		Content: req.Article.Content,
		CTime:   time.Now(),
		UTime:   time.Now(),
	}

	rpcx.CtxSet(ctx, "article", article)

	if err := s.mysqlCli.Create(article).Error; err != nil {
		return nil, errors.Wrap(err, "mysql create failed")
	}

	return nil, nil
}
