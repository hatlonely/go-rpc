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
	collection := mongoCli.Database(options.Mongo.Database).Collection(options.Mongo.Collection)
	mongoCtx, cancel := context.WithTimeout(context.Background(), options.Mongo.Timeout)
	defer cancel()
	if _, err := collection.Indexes().CreateMany(mongoCtx, []mongo.IndexModel{
		{Keys: bson.M{"name": 1}, Options: mopt.Index().SetUnique(true).SetName("nameIdx")},
	}); err != nil {
		return nil, errors.Wrap(err, "mongo.Indexes.CreateMany failed")
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
	Mongo struct {
		Database   string
		Collection string
		Timeout    time.Duration
	}
}
