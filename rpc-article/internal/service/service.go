package service

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hatlonely/go-kit/rpcx"
	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/hatlonely/go-rpc/rpc-article/api/gen/go/api"
	"github.com/hatlonely/go-rpc/rpc-article/internal/storage"
)

type ArticleService struct {
	mysqlCli *gorm.DB
	esCli    *elastic.Client
}

func NewArticleService(mysqlCli *gorm.DB, esCli *elastic.Client) (*ArticleService, error) {
	return &ArticleService{
		mysqlCli: mysqlCli,
		esCli:    esCli,
	}, nil
}

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

func (s *ArticleService) DelArticle(ctx context.Context, req *api.DelArticleReq) (*empty.Empty, error) {
	if err := s.mysqlCli.Delete(&storage.Article{}, req.Id).Error; err != nil {
		return nil, errors.Wrapf(err, "mysql delete article [%v] failed", req.Id)
	}

	return &empty.Empty{}, nil
}

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
		return nil, errors.Wrap(err, "mysql update failed")
	}

	return nil, nil
}

func (s *ArticleService) SearchArticle(ctx context.Context, req *api.SearchArticleReq) (*api.SearchArticleRes, error) {
	return nil, nil
}
