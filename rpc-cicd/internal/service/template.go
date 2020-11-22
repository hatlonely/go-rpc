package service

import (
	"context"

	"github.com/hatlonely/go-kit/rpcx"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"

	"github.com/hatlonely/go-rpc/rpc-cicd/api/gen/go/api"
)

func (s *CICDService) GetTemplate(ctx context.Context, req *api.GetTemplateReq) (*api.Template, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.TemplateCollection)
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, rpcx.NewErrorWithoutRefer(err, codes.InvalidArgument, rpcx.MetaDataGetRequestID(ctx), "InvalidObjectID", "object id is not valid")
	}
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	var template api.Template
	if err := collection.FindOne(mongoCtx, bson.M{"_id": objectID}).Decode(&template); err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.FindOne failed")
	}
	return &template, nil
}

func (s *CICDService) DelTemplate(ctx context.Context, req *api.DelTemplateReq) (*api.Empty, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.TemplateCollection)
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
	return &api.Empty{}, nil
}

func (s *CICDService) PutTemplate(ctx context.Context, req *api.PutTemplateReq) (*api.Empty, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.TemplateCollection)
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	res, err := collection.InsertOne(mongoCtx, req.Template)
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.InsertOne failed")
	}
	rpcx.CtxSet(ctx, "InsertOneRes", res)
	return &api.Empty{}, nil
}

func (s *CICDService) UpdateTemplate(ctx context.Context, req *api.UpdateTemplateReq) (*api.Empty, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.TemplateCollection)
	objectID, err := primitive.ObjectIDFromHex(req.Template.Id)
	if err != nil {
		return nil, rpcx.NewErrorWithoutRefer(err, codes.InvalidArgument, rpcx.MetaDataGetRequestID(ctx), "InvalidObjectID", "object id is not valid")
	}
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	req.Template.Id = ""
	res, err := collection.UpdateOne(mongoCtx, bson.M{"_id": objectID}, bson.M{"$set": req.Template})
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.UpdateOne failed")
	}
	rpcx.CtxSet(ctx, "UpdateOneRes", res)
	return &api.Empty{}, nil
}

func (s *CICDService) ListTemplate(ctx context.Context, req *api.ListTemplateReq) (*api.ListTemplateRes, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.TemplateCollection)
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	res, err := collection.Find(mongoCtx, bson.M{}, mopt.Find().SetLimit(req.Limit).SetSkip(req.Offset))
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.Find failed")
	}
	var templates []*api.Template
	if err := res.All(mongoCtx, &templates); err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.Find.All failed")
	}
	return &api.ListTemplateRes{Templates: templates}, nil
}