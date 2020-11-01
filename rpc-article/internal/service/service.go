package service

import (
	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic/v7"
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
