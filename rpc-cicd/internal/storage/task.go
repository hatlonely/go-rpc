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

func (s *CICDStorage) GetTask(ctx context.Context, id string) (*api.Task, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.TaskCollection)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, rpcx.NewError(codes.InvalidArgument, "InvalidObjectID", "object id is not valid", err)
	}
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	var task api.Task
	if err := collection.FindOne(mongoCtx, bson.M{"_id": objectID}).Decode(&task); err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.FindOne failed")
	}
	return &task, nil
}

func (s *CICDStorage) DelTask(ctx context.Context, id string) error {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.TaskCollection)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return rpcx.NewError(codes.InvalidArgument, "InvalidObjectID", "object id is not valid", err)
	}
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	res, err := collection.DeleteOne(mongoCtx, bson.M{"_id": objectID})
	if err != nil {
		return errors.Wrap(err, "mongo.Collection.DeleteOne failed")
	}
	rpcx.CtxSet(ctx, "DeleteOneRes", res)
	return nil
}

func (s *CICDStorage) PutTask(ctx context.Context, task *api.Task) (string, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.TaskCollection)
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	task.CreateAt = int32(time.Now().Unix())
	res, err := collection.InsertOne(mongoCtx, task)
	if err != nil {
		return "", errors.Wrap(err, "mongo.Collection.InsertOne failed")
	}
	rpcx.CtxSet(ctx, "InsertOneRes", res)
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (s *CICDStorage) UpdateTask(ctx context.Context, task *api.Task) error {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.TaskCollection)
	objectID, err := primitive.ObjectIDFromHex(task.Id)
	if err != nil {
		return rpcx.NewError(codes.InvalidArgument, "InvalidObjectID", "object id is not valid", err)
	}
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	id := task.Id
	task.Id = ""
	defer func() {
		task.Id = id
	}()
	task.UpdateAt = int32(time.Now().Unix())
	res, err := collection.UpdateOne(mongoCtx, bson.M{"_id": objectID}, bson.M{"$set": task})
	if err != nil {
		return errors.Wrap(err, "mongo.Collection.UpdateOne failed")
	}
	rpcx.CtxSet(ctx, "UpdateOneRes", res)
	return nil
}

func (s *CICDStorage) ListTask(ctx context.Context, offset int64, limit int64) ([]*api.Task, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.TaskCollection)
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	res, err := collection.Find(mongoCtx, bson.M{}, mopt.Find().SetLimit(limit).SetSkip(offset))
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.Find failed")
	}
	defer res.Close(ctx)
	var tasks []*api.Task
	if err := res.All(mongoCtx, &tasks); err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.Find.All failed")
	}
	return tasks, nil
}
