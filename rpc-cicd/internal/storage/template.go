package storage

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

func (s *CICDStorage) GetTemplate(ctx context.Context, id string) (*api.Template, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.TemplateCollection)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, rpcx.NewError(codes.InvalidArgument, "InvalidObjectID", "object id is not valid", err)
	}
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	var template api.Template
	if err := collection.FindOne(mongoCtx, bson.M{"_id": objectID}).Decode(&template); err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.FindOne failed")
	}
	return &template, nil
}

func (s *CICDStorage) DelTemplate(ctx context.Context, id string) error {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.TemplateCollection)
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

func (s *CICDStorage) PutTemplate(ctx context.Context, template *api.Template) error {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.TemplateCollection)
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	res, err := collection.InsertOne(mongoCtx, template)
	if err != nil {
		return errors.Wrap(err, "mongo.Collection.InsertOne failed")
	}
	rpcx.CtxSet(ctx, "InsertOneRes", res)
	return nil
}

func (s *CICDStorage) UpdateTemplate(ctx context.Context, template *api.Template) error {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.TemplateCollection)
	objectID, err := primitive.ObjectIDFromHex(template.Id)
	if err != nil {
		return rpcx.NewError(codes.InvalidArgument, "InvalidObjectID", "object id is not valid", err)
	}
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	template.Id = ""
	res, err := collection.UpdateOne(mongoCtx, bson.M{"_id": objectID}, bson.M{"$set": template})
	if err != nil {
		return errors.Wrap(err, "mongo.Collection.UpdateOne failed")
	}
	rpcx.CtxSet(ctx, "UpdateOneRes", res)
	return nil
}

func (s *CICDStorage) ListTemplate(ctx context.Context, offset int64, limit int64) ([]*api.Template, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.TemplateCollection)
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	res, err := collection.Find(mongoCtx, bson.M{}, mopt.Find().SetLimit(limit).SetSkip(offset))
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.Find failed")
	}
	var templates []*api.Template
	if err := res.All(mongoCtx, &templates); err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.Find.All failed")
	}
	return templates, nil
}
