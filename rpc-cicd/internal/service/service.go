package service

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
)

func NewCICDServiceWithOptions(mongoCli *mongo.Client, options *Options) (*CICDService, error) {
	{

		collection := mongoCli.Database(options.Database).Collection(options.TaskCollection)
		mongoCtx, cancel := context.WithTimeout(context.Background(), options.Timeout)
		defer cancel()
		if _, err := collection.Indexes().CreateMany(mongoCtx, []mongo.IndexModel{
			{Keys: bson.M{"name": 1}, Options: mopt.Index().SetUnique(true).SetName("nameIdx")},
		}); err != nil {
			return nil, errors.Wrap(err, "mongo.Indexes.CreateMany failed")
		}
	}
	{
		collection := mongoCli.Database(options.Database).Collection(options.VariableCollection)
		mongoCtx, cancel := context.WithTimeout(context.Background(), options.Timeout)
		defer cancel()
		if _, err := collection.Indexes().CreateMany(mongoCtx, []mongo.IndexModel{
			{Keys: bson.M{"name": 1}, Options: mopt.Index().SetUnique(true).SetName("nameIdx")},
		}); err != nil {
			return nil, errors.Wrap(err, "mongo.Indexes.CreateMany failed")
		}
	}
	{
		collection := mongoCli.Database(options.Database).Collection(options.TemplateCollection)
		mongoCtx, cancel := context.WithTimeout(context.Background(), options.Timeout)
		defer cancel()
		if _, err := collection.Indexes().CreateMany(mongoCtx, []mongo.IndexModel{
			{Keys: bson.M{"name": 1}, Options: mopt.Index().SetUnique(true).SetName("nameIdx")},
		}); err != nil {
			return nil, errors.Wrap(err, "mongo.Indexes.CreateMany failed")
		}
	}

	return &CICDService{
		mongoCli: mongoCli,
		options:  options,
	}, nil
}

type CICDService struct {
	mongoCli *mongo.Client

	options *Options
}

type Options struct {
	Database           string
	TaskCollection     string        `dft:"task"`
	JobCollection      string        `dft:"job"`
	TemplateCollection string        `dft:"template"`
	VariableCollection string        `dft:"variable"`
	Timeout            time.Duration `dft:"1s"`
}
