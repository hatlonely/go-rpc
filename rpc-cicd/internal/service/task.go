package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hatlonely/go-kit/rpcx"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"

	"github.com/hatlonely/go-rpc/rpc-cicd/api/gen/go/api"
)

func (s *CICDService) RunTask(ctx context.Context, req *api.RunTaskReq) (*api.RunTaskRes, error) {
	return &api.RunTaskRes{}, nil
}

func (s *CICDService) GetTask(ctx context.Context, req *api.GetTaskReq) (*api.Task, error) {
	collection := s.mongoCli.Database(s.options.Mongo.Database).Collection(s.options.Mongo.Collection)
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, rpcx.NewErrorWithoutRefer(err, codes.InvalidArgument, rpcx.MetaDataGetRequestID(ctx), "InvalidObjectID", "object id is not valid")
	}
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Mongo.Timeout)
	defer cancel()
	var task api.Task
	if err := collection.FindOne(mongoCtx, bson.M{"_id": objectID}).Decode(&task); err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.FindOne failed")
	}
	return &task, nil
}

func (s *CICDService) DelTask(ctx context.Context, req *api.DelTaskReq) (*empty.Empty, error) {
	collection := s.mongoCli.Database(s.options.Mongo.Database).Collection(s.options.Mongo.Collection)
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, rpcx.NewErrorWithoutRefer(err, codes.InvalidArgument, rpcx.MetaDataGetRequestID(ctx), "InvalidObjectID", "object id is not valid")
	}
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Mongo.Timeout)
	defer cancel()
	res, err := collection.DeleteOne(mongoCtx, bson.M{"_id": objectID})
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.DeleteOne failed")
	}
	rpcx.CtxSet(ctx, "DeleteOneRes", res)
	return &empty.Empty{}, nil
}

func (s *CICDService) PutTask(ctx context.Context, req *api.PutTaskReq) (*empty.Empty, error) {
	collection := s.mongoCli.Database(s.options.Mongo.Database).Collection(s.options.Mongo.Collection)
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Mongo.Timeout)
	defer cancel()
	res, err := collection.InsertOne(mongoCtx, req.Task)
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.InsertOne failed")
	}
	rpcx.CtxSet(ctx, "InsertOneRes", res)
	return &empty.Empty{}, nil
}

func (s *CICDService) UpdateTask(ctx context.Context, req *api.UpdateTaskReq) (*empty.Empty, error) {
	collection := s.mongoCli.Database(s.options.Mongo.Database).Collection(s.options.Mongo.Collection)
	objectID, err := primitive.ObjectIDFromHex(req.Task.Id)
	if err != nil {
		return nil, rpcx.NewErrorWithoutRefer(err, codes.InvalidArgument, rpcx.MetaDataGetRequestID(ctx), "InvalidObjectID", "object id is not valid")
	}
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Mongo.Timeout)
	defer cancel()
	req.Task.Id = ""
	res, err := collection.UpdateOne(mongoCtx, bson.M{"_id": objectID}, bson.M{"$set": req.Task})
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.UpdateOne failed")
	}
	rpcx.CtxSet(ctx, "UpdateOneRes", res)
	return &empty.Empty{}, nil
}

func (s *CICDService) ListTask(ctx context.Context, req *api.ListTaskReq) (*api.ListTaskRes, error) {
	collection := s.mongoCli.Database(s.options.Mongo.Database).Collection(s.options.Mongo.Collection)
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Mongo.Timeout)
	defer cancel()
	res, err := collection.Find(mongoCtx, bson.M{}, mopt.Find().SetLimit(req.Limit).SetSkip(req.Offset))
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.Find failed")
	}
	var tasks []*api.Task
	if err := res.All(mongoCtx, &tasks); err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.Find.All failed")
	}
	return &api.ListTaskRes{Tasks: tasks}, nil
}
