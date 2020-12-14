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

func (s *CICDStorage) GetVariable(ctx context.Context, id string) (*api.Variable, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.VariableCollection)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, rpcx.NewError(codes.InvalidArgument, "InvalidObjectID", "object id is not valid", err)
	}
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	var variable api.Variable
	if err := collection.FindOne(mongoCtx, bson.M{"_id": objectID}).Decode(&variable); err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.FindOne failed")
	}
	return &variable, nil
}

func (s *CICDStorage) DelVariable(ctx context.Context, id string) error {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.VariableCollection)
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

func (s *CICDStorage) PutVariable(ctx context.Context, variable *api.Variable) (string, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.VariableCollection)
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	variable.CreateAt = int32(time.Now().Unix())
	res, err := collection.InsertOne(mongoCtx, variable)
	if err != nil {
		return "", errors.Wrap(err, "mongo.Collection.InsertOne failed")
	}
	rpcx.CtxSet(ctx, "InsertOneRes", res)
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (s *CICDStorage) UpdateVariable(ctx context.Context, variable *api.Variable) error {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.VariableCollection)
	objectID, err := primitive.ObjectIDFromHex(variable.Id)
	if err != nil {
		return rpcx.NewError(codes.InvalidArgument, "InvalidObjectID", "object id is not valid", err)
	}
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	id := variable.Id
	variable.Id = ""
	defer func() {
		variable.Id = id
	}()
	variable.UpdateAt = int32(time.Now().Unix())
	res, err := collection.UpdateOne(mongoCtx, bson.M{"_id": objectID}, bson.M{"$set": variable})
	if err != nil {
		return errors.Wrap(err, "mongo.Collection.UpdateOne failed")
	}
	rpcx.CtxSet(ctx, "UpdateOneRes", res)
	return nil
}

func (s *CICDStorage) ListVariable(ctx context.Context, offset int64, limit int64) ([]*api.Variable, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.VariableCollection)
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	res, err := collection.Find(mongoCtx, bson.M{}, mopt.Find().SetLimit(limit).SetSkip(offset))
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.Find failed")
	}
	defer res.Close(ctx)
	var variables []*api.Variable
	if err := res.All(mongoCtx, &variables); err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.Find.All failed")
	}
	return variables, nil
}

func (s *CICDStorage) GetVariableByIDs(ctx context.Context, ids []string) ([]*api.Variable, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.VariableCollection)

	var objectIDs []primitive.ObjectID
	for _, i := range ids {
		objectID, err := primitive.ObjectIDFromHex(i)
		if err != nil {
			return nil, rpcx.NewError(codes.InvalidArgument, "InvalidObjectID", "object id is not valid", err)
		}
		objectIDs = append(objectIDs, objectID)
	}

	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	res, err := collection.Find(mongoCtx, bson.M{"_id": bson.M{"$in": objectIDs}})
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.Find failed")
	}
	var variables []*api.Variable
	if err := res.All(mongoCtx, &variables); err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.Find.All failed")
	}

	return variables, nil
}
