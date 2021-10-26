package main

import (
	"os"
	"testing"

	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter-cli/src/test/go/mod"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/contexts"
	"github.com/bitwormhole/starter/tests"
)

func TestCmdHelp(t *testing.T) {
	tryTestCommand(t, "help")
}

func TestCmdHelp2(t *testing.T) {
	tryTestCommand(t, "help ls")
}

func TestCmdPWD(t *testing.T) {
	tryTestCommand(t, "pwd")
}

func TestCmdLS(t *testing.T) {
	tryTestCommand(t, "ls")
}

func TestCmdCD(t *testing.T) {
	tryTestCommand(t, "cd ./../../abc")
}

func TestCmdSleep(t *testing.T) {
	tryTestCommand(t, "sleep 1000")
}

////////////////////////////////////////////////////////////////////////////////

func tryTestCommand(t *testing.T, cmd string) {
	tester := cmdTester{t: t}
	err := tester.run(cmd)
	if err != nil {
		t.Error(err)
	}
}

////////////////////////////////////////////////////////////////////////////////

type cmdTester struct {
	command string
	t       *testing.T
}

func (inst *cmdTester) run(cmd string) error {

	i := tests.Starter(inst.t)
	i.Use(mod.ExportModuleForCLITest())

	rt, err := i.RunEx()
	if err != nil {
		return err
	}

	app := rt.Context()
	contexts.SetupApplicationContext(app)
	inst.initConsole(app)

	o1, err := app.GetComponent("#cli-client-factory")
	if err != nil {
		return err
	}

	factory := o1.(cli.ClientFactory)
	client := factory.CreateClient(app)
	err = client.Execute(cmd, nil)
	if err != nil {
		return err
	}

	return rt.Exit()

}

func (inst *cmdTester) initConsole(app application.Context) {

	//	contexts.GetContextSetter(app)

	cli.SetupConsole(app, nil)
	console, err := cli.GetConsole(app)
	if err != nil {
		panic(err)
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	console.SetOutput(os.Stdout)
	console.SetError(os.Stderr)
	console.SetWD(wd)
}
