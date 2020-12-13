package service

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hatlonely/go-kit/rpcx"
	"github.com/hatlonely/go-kit/strx"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/hatlonely/go-rpc/rpc-cicd/api/gen/go/api"
	"github.com/hatlonely/go-rpc/rpc-cicd/internal/cli"
	"github.com/hatlonely/go-rpc/rpc-cicd/internal/storage"
)

var svc *CICDService

func init() {
	mongoCli, err := cli.NewMongoWithOptions(&cli.MongoOptions{
		URI:            "mongodb://localhost:27017",
		ConnectTimeout: 3 * time.Second,
		PingTimeout:    3 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	store, err := storage.NewCICDStorageWithOptions(mongoCli, &storage.Options{
		Database:           "hatlonely",
		TaskCollection:     "task",
		TemplateCollection: "template",
		VariableCollection: "variable",
		JobCollection:      "job",
		Timeout:            3 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	svc, err = NewCICDServiceWithOptions(store, &Options{
		Data: "data",
	})
	if err != nil {
		panic(err)
	}
}

func TestCICDService_PutTask(t *testing.T) {
	Convey("TestCICDService_PutTask", t, func() {
		ctx := rpcx.NewRPCXContext(context.Background())

		_, err := svc.PutTask(ctx, &api.PutTaskReq{
			Task: &api.Task{
				Name: "test-task2",
			},
		})
		So(err, ShouldBeNil)
	})
}

func TestCICDService_GetTask(t *testing.T) {
	Convey("TestCICDService_GetTask", t, func() {
		ctx := rpcx.NewRPCXContext(context.Background())

		task, err := svc.GetTask(ctx, &api.GetTaskReq{
			Id: "5fb554cc4c94d887ac699119",
		})
		So(err, ShouldBeNil)

		fmt.Println(strx.JsonMarshalIndentSortKeys(task))
	})
}

func TestCICDService_DelTask(t *testing.T) {
	Convey("TestCICDService_DelTask", t, func() {
		ctx := rpcx.NewRPCXContext(context.Background())

		_, err := svc.DelTask(ctx, &api.DelTaskReq{
			Id: "5fb554e8276455eca516e46d",
		})
		So(err, ShouldBeNil)
	})
}

func TestCICDService_UpdateTask(t *testing.T) {
	Convey("TestCICDService_UpdateTask", t, func() {
		ctx := rpcx.NewRPCXContext(context.Background())

		_, err := svc.UpdateTask(ctx, &api.UpdateTaskReq{
			Task: &api.Task{
				Id:   "5fb554cc4c94d887ac699119",
				Name: "test-task",
			},
		})
		So(err, ShouldBeNil)
	})
}

func TestCICDService_ListTask(t *testing.T) {
	Convey("TestCICDService_ListTask", t, func() {
		ctx := rpcx.NewRPCXContext(context.Background())

		tasks, err := svc.ListTask(ctx, &api.ListTaskReq{
			Offset: 1,
			Limit:  2,
		})
		So(err, ShouldBeNil)

		fmt.Println(strx.JsonMarshalIndentSortKeys(tasks))
	})
}
