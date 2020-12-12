package storage

import (
	"context"
	"time"

	"github.com/hatlonely/go-kit/rpcx"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"

	"github.com/hatlonely/go-rpc/rpc-cicd/api/gen/go/api"
)

func (s *CICDStorage) GetJob(ctx context.Context, id string) (*api.Job, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.JobCollection)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, rpcx.NewError(codes.InvalidArgument, "InvalidObjectID", "object id is not valid", err)
	}
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	var job api.Job
	if err := collection.FindOne(mongoCtx, bson.M{"_id": objectID}).Decode(&job); err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.FindOne failed")
	}
	return &job, nil
}

func (s *CICDStorage) DelJob(ctx context.Context, id string) error {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.JobCollection)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return rpcx.NewError(codes.InvalidArgument, "InvalidObjectID", "object id is not valid", err)
	}
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	if _, err := collection.DeleteOne(mongoCtx, bson.M{"_id": objectID}); err != nil {
		return errors.Wrap(err, "mongo.Collection.DeleteOne failed")
	}
	return nil
}

func (s *CICDStorage) PutJob(ctx context.Context, job *api.Job) (string, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.JobCollection)
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	seq, err := s.Inc(ctx, job.TaskID)
	if err != nil {
		return "", errors.Wrap(err, "CICDStorage.Inc failed")
	}
	job.Seq = seq
	job.CreateAt = int32(time.Now().Unix())
	res, err := collection.InsertOne(mongoCtx, job)
	if err != nil {
		return "", errors.Wrap(err, "mongo.Collection.InsertOne failed")
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (s *CICDStorage) UpdateJob(ctx context.Context, job *api.Job) error {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.JobCollection)
	objectID, err := primitive.ObjectIDFromHex(job.Id)
	if err != nil {
		return rpcx.NewError(codes.InvalidArgument, "InvalidObjectID", "object id is not valid", err)
	}
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	id := job.Id
	job.Id = ""
	defer func() {
		job.Id = id
	}()
	job.UpdateAt = int32(time.Now().Unix())
	if _, err := collection.UpdateOne(mongoCtx, bson.M{"_id": objectID}, bson.M{"$set": job}); err != nil {
		return errors.Wrap(err, "mongo.Collection.UpdateOne failed")
	}
	return nil
}

func (s *CICDStorage) ListJob(ctx context.Context, taskID string, offset int64, limit int64) ([]*api.Job, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.JobCollection)
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	res, err := collection.Find(mongoCtx, bson.M{"taskID": taskID}, mopt.Find().SetLimit(limit).SetSkip(offset).SetSort(bson.M{"createAt": -1}))
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.Find failed")
	}
	var jobs []*api.Job
	if err := res.All(mongoCtx, &jobs); err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.Find.All failed")
	}
	return jobs, nil
}
