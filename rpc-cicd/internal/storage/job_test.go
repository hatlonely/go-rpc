package storage

import (
	"context"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/hatlonely/go-rpc/rpc-cicd/api/gen/go/api"
	"github.com/hatlonely/go-rpc/rpc-cicd/internal/cli"
)

func TestCICDStorage(t *testing.T) {
	Convey("CICDStorage", t, func() {
		mongoCli, err := cli.NewMongoWithOptions(&cli.MongoOptions{
			URI:            "mongodb://localhost:27017",
			ConnectTimeout: 3 * time.Second,
			PingTimeout:    2 * time.Second,
		})
		So(err, ShouldBeNil)
		store, err := NewCICDStorageWithOptions(mongoCli, &Options{
			Database:           "test",
			TaskCollection:     "task",
			JobCollection:      "job",
			TemplateCollection: "template",
			VariableCollection: "variable",
			SequenceCollection: "sequence",
			Timeout:            time.Second,
			JobExpiration:      72 * time.Hour,
		})
		So(err, ShouldBeNil)

		Convey("Job", func() {
			_, _ = mongoCli.Database("test").Collection("job").DeleteMany(context.Background(), bson.M{"taskName": "test-task"})
			jobID, err := store.PutJob(context.Background(), &api.Job{
				TaskName: "test-task",
				Status:   JobStatusWaiting,
			})
			So(err, ShouldBeNil)
			So(jobID, ShouldNotBeEmpty)

			job, err := store.GetJob(context.Background(), jobID)
			So(err, ShouldBeNil)
			So(job.TaskName, ShouldEqual, "test-task")
			So(job.Status, ShouldEqual, JobStatusWaiting)

			job.Status = JobStatusRunning
			So(store.UpdateJob(context.Background(), job), ShouldBeNil)
			job1, err := store.GetJob(context.Background(), jobID)
			So(err, ShouldBeNil)
			So(job, ShouldResemble, job1)

			So(store.DelJob(context.Background(), jobID), ShouldBeNil)
		})

		Convey("find Job", func() {
			_, _ = mongoCli.Database("test").Collection("job").DeleteMany(context.Background(), bson.M{"taskName": "test-task"})
			jobID, err := store.PutJob(context.Background(), &api.Job{
				TaskName: "test-task",
				Status:   JobStatusWaiting,
			})
			So(err, ShouldBeNil)
			So(jobID, ShouldNotBeEmpty)

			{
				job, err := store.FindOneUnscheduleJob(context.Background())
				So(err, ShouldBeNil)
				So(job.Id, ShouldEqual, jobID)
			}
			{
				ok, err := store.UpdateJobStatus(context.Background(), jobID, JobStatusWaiting, JobStatusRunning)
				So(err, ShouldBeNil)
				So(ok, ShouldBeTrue)
			}
			{
				ok, err := store.UpdateJobStatus(context.Background(), jobID, JobStatusWaiting, JobStatusRunning)
				So(err, ShouldBeNil)
				So(ok, ShouldBeFalse)
			}
			{
				job, err := store.FindOneUnscheduleJob(context.Background())
				So(err, ShouldBeNil)
				So(job, ShouldBeNil)
			}

			So(store.DelJob(context.Background(), jobID), ShouldBeNil)
		})

		Convey("Task", func() {
			_, _ = mongoCli.Database("test").Collection("task").DeleteOne(context.Background(), bson.M{"name": "test-task"})
			taskID, err := store.PutTask(context.Background(), &api.Task{
				Name:        "test-task",
				Description: "test-description",
				TemplateIDs: []string{"tpl1", "tpl2"},
				VariableIDs: []string{"var1", "var2"},
			})
			So(err, ShouldBeNil)
			So(taskID, ShouldNotBeEmpty)

			task, err := store.GetTask(context.Background(), taskID)
			So(err, ShouldBeNil)
			So(task.Name, ShouldEqual, "test-task")
			So(task.Description, ShouldEqual, "test-description")
			So(task.TemplateIDs, ShouldResemble, []string{"tpl1", "tpl2"})
			So(task.VariableIDs, ShouldResemble, []string{"var1", "var2"})

			task.Name = "test-task1"
			task.Description = "test-description1"
			task.TemplateIDs = []string{"tpl3", "tpl4", "tpl5"}
			task.VariableIDs = []string{"var3"}
			So(store.UpdateTask(context.Background(), task), ShouldBeNil)
			task1, err := store.GetTask(context.Background(), taskID)
			So(err, ShouldBeNil)
			So(task, ShouldResemble, task1)

			So(store.DelTask(context.Background(), taskID), ShouldBeNil)
		})
	})
}
