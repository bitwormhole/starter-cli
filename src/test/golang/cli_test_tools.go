package golang

import (
	"context"
	"testing"

	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter-cli/src/test/golang/mod"
	"github.com/bitwormhole/starter-cli/terminal"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/tests"
)

// OnTest CLI 测试的回调函数
type OnTest func(tool *CLITesting) error

// RunTest 以回调的方式运行测试
func RunTest(t *testing.T, h OnTest) {
	tester := cmdTester{}
	tester.t = t
	tester.callback = h
	err := tester.run()
	if err != nil {
		t.Error(err)
	}
}

// CLITesting  回调的环境参数
type CLITesting struct {
	t       *testing.T
	appCtx  application.Context
	ctx     context.Context
	client  cli.Client
	factory cli.ClientFactory
}

////////////////////////////////////////////////////////////////////////////////

type cmdTester struct {
	//  command  string
	t        *testing.T
	callback OnTest
}

func (inst *cmdTester) run() error {

	i := tests.Starter(inst.t)
	i.Use(mod.ExportModuleForCLITest())

	rt, err := i.RunEx()
	if err != nil {
		return err
	}

	ctx1 := rt.Context()
	ctx2, err := inst.prepareCLIContext(ctx1)
	if err != nil {
		return err
	}

	o1, err := ctx1.GetComponent("#cli-client-factory")
	if err != nil {
		return err
	}

	factory := o1.(cli.ClientFactory)
	client := factory.CreateClient(ctx2)

	callback := &CLITesting{}
	callback.client = client
	callback.factory = factory
	callback.appCtx = ctx1
	callback.ctx = ctx2
	callback.t = inst.t

	err = inst.callback(callback)
	if err != nil {
		return err
	}

	return rt.Exit()
}

func (inst *cmdTester) prepareCLIContext(app application.Context) (context.Context, error) {
	return terminal.Prepare(app)
}
