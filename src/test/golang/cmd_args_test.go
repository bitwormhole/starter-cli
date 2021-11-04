package golang

import (
	"testing"

	"github.com/bitwormhole/starter-cli/src/test/golang/cmds"
)

func TestCLIArguments1(t *testing.T) {

	const CMD = cmds.CommandArgsBypass
	cat := cliArgumentsTester{t: t}

	cat.inputCommand = CMD + "  a b  C  D "
	cat.inputArgs = nil

	cat.wantCommand = CMD
	cat.wantArgs = []string{CMD, "a", "b", "C", "D"}

	cat.run()
}

func TestCLIArguments2(t *testing.T) {

	const CMD = cmds.CommandArgsBypass
	cat := cliArgumentsTester{t: t}

	cat.inputCommand = ""
	cat.inputArgs = []string{CMD, "a", "b", "C", "D"}

	cat.wantCommand = CMD
	cat.wantArgs = []string{CMD, "a", "b", "C", "D"}

	cat.run()
}

func TestCLIArguments3(t *testing.T) {

	const CMD = cmds.CommandArgsBypass
	cat := cliArgumentsTester{t: t}

	cat.inputCommand = CMD + "   a b"
	cat.inputArgs = []string{"C", "D"}

	cat.wantCommand = CMD
	cat.wantArgs = []string{CMD, "a", "b", "C", "D"}

	cat.run()
}

func TestCLIArguments4(t *testing.T) {

	const CMD = cmds.CommandArgsBypass
	cat := cliArgumentsTester{t: t}

	cat.inputCommand = CMD
	cat.inputArgs = []string{"C", "D"}

	cat.wantCommand = CMD
	cat.wantArgs = []string{CMD, "C", "D"}

	cat.run()
}

////////////////////////////////////////////////////////////////////////////////

type cliArgumentsTester struct {
	t *testing.T

	inputCommand string
	inputArgs    []string

	wantCommand string
	wantArgs    []string
}

func (inst *cliArgumentsTester) run() {
	RunTest(inst.t, func(tool *CLITesting) error {
		bypass := cmds.GetBypassArgs(tool.ctx)
		bypass.WantCmd = inst.wantCommand
		bypass.WantArgs = inst.wantArgs
		return tool.client.Execute(inst.inputCommand, inst.inputArgs)
	})
}
