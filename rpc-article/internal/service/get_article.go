package service

import (
	"context"
	"time"

	"github.com/hatlonely/go-kit/rpcx"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/hatlonely/go-rpc/rpc-article/api/gen/go/api"
	"github.com/hatlonely/go-rpc/rpc-article/internal/storage"
)

func (s *ArticleService) GetArticle(ctx context.Context, req *api.GetArticleReq) (*api.Article, error) {
	requestID := rpcx.MetaDataGetRequestID(ctx)

	article := &storage.Article{}
	if err := s.mysqlCli.Where("id=?", req.Id).First(article).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, rpcx.NewErrorWithoutReferf(codes.NotFound, requestID, "NotFound", "article [%v] not exist", req.Id)
		}
		return nil, errors.Wrapf(err, "mysql select article [%v] failed", req.Id)
	}

	rpcx.CtxSet(ctx, "article", article)

	return &api.Article{
		Id:      int64(article.ID),
		OwnerID: int64(article.OwnerID),
		Title:   article.Title,
		Tags:    article.Tags,
		Brief:   article.Brief,
		Content: article.Content,
		CTime:   article.CTime.Format(time.RFC3339),
		UTime:   article.CTime.Format(time.RFC3339),
	}, nil
}
