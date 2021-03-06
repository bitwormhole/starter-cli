package support

import (
	"context"

	"github.com/bitwormhole/starter-cli/cli"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/task"
)

// DefaultClientFactory 命令客户端工厂
type DefaultClientFactory struct {
	markup.Component `id:"cli-client-factory"`

	CLI *Context `inject:"#cli-context"`
}

func (inst *DefaultClientFactory) _Impl() cli.ClientFactory {
	return inst
}

// CreateClient 创建同步客户端
func (inst *DefaultClientFactory) CreateClient(ctx context.Context) cli.Client {
	service := inst.CLI.Service
	client := &syncClientImpl{
		context: ctx,
		factory: inst,
		service: service,
	}
	return client
}

// CreateAsyncClient 创建异步客户端
func (inst *DefaultClientFactory) CreateAsyncClient(ctx context.Context) cli.AsyncClient {
	service := inst.CLI.Service
	client1 := &syncClientImpl{
		context: ctx,
		factory: inst,
		service: service,
	}
	client2 := &asyncClientImpl{
		sync: client1,
	}
	return client2
}

////////////////////////////////////////////////////////////////////////////////

// 执行命令的客户端
type syncClientImpl struct {
	context context.Context
	factory cli.ClientFactory
	service cli.Service
}

func (inst *syncClientImpl) _Impl() cli.Client {
	return inst
}

func (inst *syncClientImpl) GetFactory() cli.ClientFactory {
	return inst.factory
}

func (inst *syncClientImpl) GetContext() context.Context {
	return inst.context
}

// Execute  todo...
func (inst *syncClientImpl) ExecuteTask(t *cli.Task) error {
	tc := inst.makeTaskContext(t)
	chain := inst.service.GetFilterChain()
	return chain.Handle(tc)
}

// ExecuteWithArguments  todo...
func (inst *syncClientImpl) Execute(cmd string, args []string) error {
	t := inst.makeTaskWithArgs(cmd, args)
	return inst.ExecuteTask(t)
}

// ExecuteScript  todo...
func (inst *syncClientImpl) ExecuteScript(script string) error {
	t := inst.makeTaskWithScript(script)
	return inst.ExecuteTask(t)
}

func (inst *syncClientImpl) makeTaskWithArgs(cmd string, args1 []string) *cli.Task {

	clbuilder := cli.CommandLineBuilder{}
	parser := cli.CommandLineParser{}

	args0, err := parser.Parse(cmd)
	if err != nil {
		args0 = []string{}
	}

	clbuilder.AppendStrings(args0)
	clbuilder.AppendStrings(args1)
	line2 := clbuilder.Create()

	args2, err := parser.Parse(line2)
	if err != nil {
		args2 = []string{}
	}

	builder := cli.TaskListBuilder{}
	builder.AddLine(line2, 0, args2)
	t := &cli.Task{}
	t.Script = line2
	t.TaskList = builder.Create()
	return t
}

func (inst *syncClientImpl) makeTaskWithScript(script string) *cli.Task {

	builder := cli.TaskListBuilder{}
	builder.ParseScript(script)

	t := &cli.Task{}
	t.Script = script
	t.TaskList = builder.Create()
	return t
}

func (inst *syncClientImpl) makeTaskContext(t *cli.Task) *cli.TaskContext {

	script := t.Script
	tasklist := t.TaskList
	context := t.Context
	reporter := t.Reporter

	if tasklist == nil {
		task2 := inst.makeTaskWithScript(script)
		tasklist = task2.TaskList
	}

	if context == nil {
		context = inst.context
	}

	if reporter == nil {
		r, err := task.GetProgressReporter(context)
		if err == nil {
			reporter = r
		} else {
			reporter = &MockReporter{}
		}
	}

	tc := &cli.TaskContext{}
	tc.Context = context
	tc.TaskList = tasklist
	tc.Reporter = reporter

	return tc
}

////////////////////////////////////////////////////////////////////////////////

type asyncClientImpl struct {
	sync *syncClientImpl
}

func (inst *asyncClientImpl) _Impl() cli.AsyncClient {
	return inst
}

func (inst *asyncClientImpl) GetContext() context.Context {
	return inst.sync.context
}

func (inst *asyncClientImpl) GetFactory() cli.ClientFactory {
	return inst.sync.factory
}

func (inst *asyncClientImpl) Execute(cmd string, args []string) task.Promise {
	t := inst.sync.makeTaskWithArgs(cmd, args)
	return inst.ExecuteTask(t)
}

func (inst *asyncClientImpl) ExecuteScript(script string) task.Promise {
	t := inst.sync.makeTaskWithScript(script)
	return inst.ExecuteTask(t)
}

func (inst *asyncClientImpl) ExecuteTask(t *cli.Task) task.Promise {

	return task.NewPromise(func(resolve task.ResolveFn, reject task.RejectFn) {
		err := inst.sync.ExecuteTask(t)
		if err == nil {
			resolve(t)
		} else {
			reject(err)
		}
	})
}
