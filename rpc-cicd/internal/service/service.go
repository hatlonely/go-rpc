package service

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hatlonely/go-rpc/rpc-cicd/internal/storage"
)

func NewCICDServiceWithOptions(storage *storage.CICDStorage, mongoCli *mongo.Client, options *Options) (*CICDService, error) {
	return &CICDService{
		options: options,
		storage: storage,
	}, nil
}

type CICDService struct {
	storage *storage.CICDStorage

	options *Options
}

type Options struct {
}
