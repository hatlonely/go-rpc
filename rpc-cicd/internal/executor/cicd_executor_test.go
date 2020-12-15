package executor

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/hatlonely/go-kit/cli"

	"github.com/hatlonely/go-rpc/rpc-cicd/internal/storage"
)

func TestCICDExecutor(t *testing.T) {
	Convey("TextExecutor", t, func() {
		mongoCli, err := cli.NewMongoWithOptions(&cli.MongoOptions{
			URI:            "mongodb://localhost:27017",
			ConnectTimeout: 3 * time.Second,
			PingTimeout:    2 * time.Second,
		})
		So(err, ShouldBeNil)
		store, err := storage.NewCICDStorageWithOptions(mongoCli, &storage.Options{
			Database:           "test",
			TaskCollection:     "task",
			JobCollection:      "job",
			TemplateCollection: "template",
			VariableCollection: "variable",
			SequenceCollection: "sequence",
			Timeout:            time.Second,
			JobExpiration:      72 * time.Hour,
		})

		executor := NewCICDExecutorWithOptions(store, &CICDOptions{
			WorkerNum: 2,
			Data:      "data",
			SleepTime: 2 * time.Second,
		})

		executor.Run()

		time.Sleep(3 * time.Second)

		executor.Stop()
	})
}
