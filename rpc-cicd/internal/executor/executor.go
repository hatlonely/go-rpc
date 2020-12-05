package executor

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/hatlonely/go-kit/logger"
)

func NewExecutorWithOptions(handler Handler, options *Options) (*Executor, error) {
	ctx, cancel := context.WithCancel(context.Background())
	return &Executor{
		taskQueue: make(chan *Task, options.QueueLen),
		options:   options,
		handler:   handler,
		ctx:       ctx,
		cancel:    cancel,
		logger:    logger.NewStdoutJsonLogger(),
	}, nil
}

type Options struct {
	QueueLen  int `dft:"200"`
	WorkerNum int `dft:"20"`
}

type Executor struct {
	wg      sync.WaitGroup
	handler Handler
	logger  *logger.Logger

	taskQueue chan *Task
	options   *Options

	ctx    context.Context
	cancel context.CancelFunc
}

type Handler func(context.Context, interface{}) error

func (e *Executor) Run() {
	for i := 0; i < e.options.WorkerNum; i++ {
		e.wg.Add(1)
		go func() {
			e.work(e.ctx)
			e.wg.Done()
		}()
	}
}

func (e *Executor) Stop() {
	close(e.taskQueue)
	e.cancel()
	e.wg.Wait()
}

type Task struct {
	id   string
	task interface{}
}

func (e *Executor) AddTask(id string, task interface{}) {
	e.taskQueue <- &Task{id: id, task: task}
}

func (e *Executor) work(ctx context.Context) {
	for task := range e.taskQueue {
		ctx = NewExecutorContext(ctx)

		ts := time.Now()
		err := e.handler(ctx, task.task)

		e.logger.Info(map[string]interface{}{
			"id":       task.id,
			"task":     task.task,
			"err":      err,
			"errStack": fmt.Sprintf("%+v", err),
			"timeMs":   time.Now().Sub(ts).Milliseconds(),
			"ctx":      ctx.Value(executorKey{}),
		})
	}
}

type executorKey struct{}

func NewExecutorContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, executorKey{}, map[string]interface{}{})
}

func CtxSet(ctx context.Context, key string, val interface{}) {
	m := ctx.Value(executorKey{}).(map[string]interface{})
	m[key] = val
}

func CtxGet(ctx context.Context, key string) interface{} {
	m := ctx.Value(executorKey{}).(map[string]interface{})
	return m[key]
}

func FromExecutorContext(ctx context.Context) map[string]interface{} {
	return ctx.Value(executorKey{}).(map[string]interface{})
}
