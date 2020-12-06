package service

import (
	"github.com/hatlonely/go-rpc/rpc-cicd/internal/executor"
	"github.com/hatlonely/go-rpc/rpc-cicd/internal/storage"
)

func NewCICDServiceWithOptions(storage *storage.CICDStorage, options *Options) (*CICDService, error) {
	svc := &CICDService{
		options: options,
		storage: storage,
	}

	return svc, nil
}

func (s *CICDService) SetExecutor(executor *executor.Executor) {
	s.executor = executor
}

type CICDService struct {
	storage  *storage.CICDStorage
	executor *executor.Executor

	options *Options
}

type Options struct {
	Data string `dft:"data"`
}
