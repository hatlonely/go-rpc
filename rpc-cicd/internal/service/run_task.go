package service

import (
	"context"
	"encoding/json"

	"github.com/hatlonely/go-kit/rpcx"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"

	"github.com/hatlonely/go-rpc/rpc-cicd/api/gen/go/api"
)

const JobStatusRunning = "Running"
const JobStatusWaiting = "Waiting"
const JobStatusFailed = "Failed"
const JobStatusFinish = "Finish"

func (s *CICDService) RunTask(ctx context.Context, req *api.RunTaskReq) (*api.RunTaskRes, error) {
	// 获取 task
	taskCollection := s.mongoCli.Database(s.options.Database).Collection(s.options.TaskCollection)
	objectID, err := primitive.ObjectIDFromHex(req.TaskID)
	if err != nil {
		return nil, rpcx.NewError(codes.InvalidArgument, "InvalidObjectID", "object id is not valid", err)
	}
	mongoCtx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	var task api.Task
	if err := taskCollection.FindOne(mongoCtx, bson.M{"_id": objectID}).Decode(&task); err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.FindOne failed")
	}

	jobID := primitive.NewObjectID()
	job := api.Job{
		Id:     jobID.String(),
		TaskID: req.TaskID,
		Status: JobStatusRunning,
	}

	// 插入 job
	jobCollection := s.mongoCli.Database(s.options.Database).Collection(s.options.JobCollection)
	mongoCtx, cancel = context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	res, err := jobCollection.InsertOne(mongoCtx, job)
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.InsertOne failed")
	}
	rpcx.CtxSet(ctx, "InsertOneRes", res)

	// 执行 job
	go s.runTask(jobID.String(), task)

	return &api.RunTaskRes{JobID: jobID.String()}, nil
}

func (s *CICDService) GetTemplates(ctx context.Context, req *api.GetTemplatesReq) (*api.ListTemplateRes, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.TemplateCollection)

	var objectIDs []primitive.ObjectID
	for _, i := range req.Ids {
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
	var templates []*api.Template
	if err := res.All(mongoCtx, &templates); err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.Find.All failed")
	}

	return &api.ListTemplateRes{Templates: templates}, nil
}

func (s *CICDService) GetVariables(ctx context.Context, req *api.GetVariablesReq) (*api.ListVariableRes, error) {
	collection := s.mongoCli.Database(s.options.Database).Collection(s.options.VariableCollection)

	var objectIDs []primitive.ObjectID
	for _, i := range req.Ids {
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

	return &api.ListVariableRes{Variables: variables}, nil
}

func mergeVariables(variables []api.Variable) (map[string]interface{}, error) {
	kvs := map[string]interface{}{}

	for _, variable := range variables {
		var v interface{}
		if err := json.Unmarshal([]byte(variable.Kvs), &v); err != nil {
			return nil, errors.Wrap(err, "mergeVariables failed")
		}

		kvs[variable.Name] = v
	}

	return kvs, nil
}

func (s *CICDService) runTask(jobID string, task api.Task) error {
	return nil
}
