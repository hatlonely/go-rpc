package executor

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExecutor(t *testing.T) {
	Convey("TextExecutor", t, func() {
		executor := NewExecutorWithOptions(func(ctx context.Context, v interface{}) error {
			CtxSet(ctx, "key1", "val1")
			CtxSet(ctx, "key2", "val2")
			return nil
		}, &Options{
			QueueLen:  20,
			WorkerNum: 20,
		})

		executor.Run()

		executor.AddTask("task1", "hello world")
		executor.AddTask("task2", "hello golang")

		executor.Stop()
	})
}
