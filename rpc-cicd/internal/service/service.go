package service

import (
	"github.com/hatlonely/go-rpc/rpc-cicd/internal/storage"
)

func NewCICDServiceWithOptions(storage *storage.CICDStorage, options *Options) (*CICDService, error) {
	svc := &CICDService{
		options: options,
		storage: storage,
	}

	return svc, nil
}

type CICDService struct {
	storage *storage.CICDStorage
	options *Options
}

type Options struct {
	Data string `dft:"data"`
}
