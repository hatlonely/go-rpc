package service

import (
	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic/v7"
)

type AncientService struct {
	mysqlCli *gorm.DB
	esCli    *elastic.Client
}

func NewAncientService(mysqlCli *gorm.DB, esCli *elastic.Client) (*AncientService, error) {
	return &AncientService{
		mysqlCli: mysqlCli,
		esCli:    esCli,
	}, nil
}
