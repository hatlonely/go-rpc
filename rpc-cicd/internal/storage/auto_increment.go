package storage

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
)

type AutoIncrement struct {
	Key string `bson:"key"`
	Val int32  `bson:"val"`
}

func (s *CICDStorage) Inc(ctx context.Context, key string) (int32, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.AutoIncrementCollection)
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	var autoIncrement AutoIncrement
	if err := collection.FindOneAndUpdate(
		mongoCtx, bson.M{"key": key}, bson.M{"$inc": bson.M{"val": 1}},
		mopt.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(mopt.After),
	).Decode(&autoIncrement); err != nil {
		return 0, errors.Wrap(err, "mongo.Collection.FindOneAndUpdate failed")
	}
	return autoIncrement.Val, nil
}
