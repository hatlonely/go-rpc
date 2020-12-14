package executor

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/hatlonely/go-kit/logger"
)

func NewExecutorWithOptions(handler Handler, options *Options) *Executor {
	ctx, cancel := context.WithCancel(context.Background())
	return &Executor{
		taskQueue: make(chan *Task, options.QueueLen),
		options:   options,
		handler:   handler,
		ctx:       ctx,
		cancel:    cancel,
		logger:    logger.NewStdoutJsonLogger(),
	}
}

func (e *Executor) SetLogger(log *logger.Logger) {
	e.logger = log
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

	cancelTaskIDs       sync.Map
	taskIDCancelFuncMap sync.Map
}

type Handler func(context.Context, interface{}) error

func (e *Executor) Run() {
	for i := 0; i < e.options.WorkerNum; i++ {
		e.wg.Add(1)
		go func() {
			e.handleTask(e.ctx)
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

func (e *Executor) CancelTask(id string) {
	if cancel, ok := e.taskIDCancelFuncMap.Load(id); ok {
		cancel.(context.CancelFunc)()
	} else {
		e.cancelTaskIDs.Store(id, struct{}{})
	}
}

func (e *Executor) handleTask(ctx context.Context) {
	for task := range e.taskQueue {
		ts := time.Now()
		var err error
		var isCancel bool
		if _, ok := e.cancelTaskIDs.Load(task.id); ok {
			e.cancelTaskIDs.Delete(task.id)
			isCancel = true
		} else {
			var cancel context.CancelFunc
			ctx, cancel = context.WithCancel(NewExecutorContext(ctx))

			e.taskIDCancelFuncMap.Store(task.id, cancel)
			err = e.handler(ctx, task.task)
			cancel()
			e.taskIDCancelFuncMap.Delete(task.id)
		}

		e.logger.Info(map[string]interface{}{
			"id":       task.id,
			"task":     task.task,
			"err":      err,
			"errStack": fmt.Sprintf("%+v", err),
			"timeMs":   time.Now().Sub(ts).Milliseconds(),
			"ctx":      ctx.Value(executorKey{}),
			"isCancel": isCancel,
		})
	}
}

type executorKey struct{}

func NewExecutorContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, executorKey{}, map[string]interface{}{})
}

func CtxSet(ctx context.Context, key string, val interface{}) {
	m := ctx.Value(executorKey{})
	if m == nil {
		return
	}
	m.(map[string]interface{})[key] = val
}

func CtxGet(ctx context.Context, key string) interface{} {
	m := ctx.Value(executorKey{})
	if m == nil {
		return nil
	}
	return m.(map[string]interface{})[key]
}

func FromExecutorContext(ctx context.Context) map[string]interface{} {
	m := ctx.Value(executorKey{})
	if m == nil {
		return nil
	}
	return m.(map[string]interface{})
}
