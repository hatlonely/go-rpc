package storage

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
)

type CICDStorage struct {
	mongoCli *mongo.Client
	options  *Options
}

type Options struct {
	Database           string
	TaskCollection     string        `dft:"task"`
	JobCollection      string        `dft:"job"`
	TemplateCollection string        `dft:"template"`
	VariableCollection string        `dft:"variable"`
	SequenceCollection string        `dft:"sequence"`
	Timeout            time.Duration `dft:"1s"`
}

func NewCICDStorageWithOptions(mongoCli *mongo.Client, options *Options) (*CICDStorage, error) {
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
	{
		collection := mongoCli.Database(options.Database).Collection(options.JobCollection)
		mongoCtx, cancel := context.WithTimeout(context.Background(), options.Timeout)
		defer cancel()
		if _, err := collection.Indexes().CreateMany(mongoCtx, []mongo.IndexModel{
			{Keys: bson.M{"taskID": 1}, Options: mopt.Index().SetName("taskIDIdx")},
			{Keys: bson.M{"createAt": 1}, Options: mopt.Index().SetExpireAfterSeconds(3600)},
		}); err != nil {
			return nil, errors.Wrap(err, "mongo.Indexes.CreateMany failed")
		}
	}
	{
		collection := mongoCli.Database(options.Database).Collection(options.SequenceCollection)
		mongoCtx, cancel := context.WithTimeout(context.Background(), options.Timeout)
		defer cancel()
		if _, err := collection.Indexes().CreateMany(mongoCtx, []mongo.IndexModel{
			{Keys: bson.M{"key": 1}, Options: mopt.Index().SetUnique(true).SetName("keyIdx")},
		}); err != nil {
			return nil, errors.Wrap(err, "mongo.Indexes.CreateMany failed")
		}
	}

	return &CICDStorage{
		mongoCli: mongoCli,
		options:  options,
	}, nil
}
