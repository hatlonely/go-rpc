package storage

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
)

type Sequence struct {
	Key string `bson:"key"`
	Val int32  `bson:"val"`
}

func (s *CICDStorage) IncSequence(ctx context.Context, key string) (int32, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.SequenceCollection)
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	var sequence Sequence
	if err := collection.FindOneAndUpdate(
		mongoCtx, bson.M{"key": key}, bson.M{"$inc": bson.M{"val": 1}},
		mopt.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(mopt.After),
	).Decode(&sequence); err != nil {
		return 0, errors.Wrap(err, "mongo.Collection.FindOneAndUpdate failed")
	}
	return sequence.Val, nil
}

func (s *CICDStorage) DelSequence(ctx context.Context, key string) error {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.SequenceCollection)
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	if _, err := collection.DeleteOne(mongoCtx, bson.M{"key": key}); err != nil {
		return errors.Wrap(err, "mongo.Collection.DeleteOne failed")
	}
	return nil
}
