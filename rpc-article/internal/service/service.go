package service

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"

	"github.com/hatlonely/go-rpc/rpc-article/internal/storage"
)

type ArticleService struct {
	mysqlCli *gorm.DB
	esCli    *elastic.Client
}

func NewArticleService(mysqlCli *gorm.DB, esCli *elastic.Client) (*ArticleService, error) {
	if !mysqlCli.HasTable(&storage.Article{}) {
		if err := mysqlCli.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&storage.Article{}).Error; err != nil {
			return nil, err
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	exists, err := esCli.IndexExists("article").Do(ctx)
	if err != nil {
		return nil, err
	}
	if !exists {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		createIndex, err := esCli.CreateIndex("ancient").Body(storage.ArticleMappingForElasticsearch).Do(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "esCli.CreateIndex failed.")
		}
		if !createIndex.Acknowledged {
			return nil, errors.New("esCli.CreateIndex not acknowledged")
		}
	}

	return &ArticleService{
		mysqlCli: mysqlCli,
		esCli:    esCli,
	}, nil
}
