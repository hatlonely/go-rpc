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

func (s *CICDService) GetVariable(ctx context.Context, req *api.GetVariableReq) (*api.Variable, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.VariableCollection)
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, rpcx.NewErrorWithoutRefer(err, codes.InvalidArgument, rpcx.MetaDataGetRequestID(ctx), "InvalidObjectID", "object id is not valid")
	}
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	var variable api.Variable
	if err := collection.FindOne(mongoCtx, bson.M{"_id": objectID}).Decode(&variable); err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.FindOne failed")
	}
	return &variable, nil
}

func (s *CICDService) DelVariable(ctx context.Context, req *api.DelVariableReq) (*empty.Empty, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.VariableCollection)
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, rpcx.NewErrorWithoutRefer(err, codes.InvalidArgument, rpcx.MetaDataGetRequestID(ctx), "InvalidObjectID", "object id is not valid")
	}
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	res, err := collection.DeleteOne(mongoCtx, bson.M{"_id": objectID})
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.DeleteOne failed")
	}
	rpcx.CtxSet(ctx, "DeleteOneRes", res)
	return &empty.Empty{}, nil
}

func (s *CICDService) PutVariable(ctx context.Context, req *api.PutVariableReq) (*empty.Empty, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.VariableCollection)
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	res, err := collection.InsertOne(mongoCtx, req.Variable)
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.InsertOne failed")
	}
	rpcx.CtxSet(ctx, "InsertOneRes", res)
	return &empty.Empty{}, nil
}

func (s *CICDService) UpdateVariable(ctx context.Context, req *api.UpdateVariableReq) (*empty.Empty, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.VariableCollection)
	objectID, err := primitive.ObjectIDFromHex(req.Variable.Id)
	if err != nil {
		return nil, rpcx.NewErrorWithoutRefer(err, codes.InvalidArgument, rpcx.MetaDataGetRequestID(ctx), "InvalidObjectID", "object id is not valid")
	}
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	req.Variable.Id = ""
	res, err := collection.UpdateOne(mongoCtx, bson.M{"_id": objectID}, bson.M{"$set": req.Variable})
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.UpdateOne failed")
	}
	rpcx.CtxSet(ctx, "UpdateOneRes", res)
	return &empty.Empty{}, nil
}

func (s *CICDService) ListVariable(ctx context.Context, req *api.ListVariableReq) (*api.ListVariableRes, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.VariableCollection)
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	res, err := collection.Find(mongoCtx, bson.M{}, mopt.Find().SetLimit(req.Limit).SetSkip(req.Offset))
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.Find failed")
	}
	var variables []*api.Variable
	if err := res.All(mongoCtx, &variables); err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.Find.All failed")
	}
	return &api.ListVariableRes{Variables: variables}, nil
}
