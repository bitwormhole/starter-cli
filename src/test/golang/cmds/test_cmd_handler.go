package cmds

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/contexts"
	"github.com/bitwormhole/starter/markup"
)

const CommandArgsBypass = "test-cli-args-bypass"

type TestCommandHandler struct {
	markup.Component `class:"cli-handler"`

	Context application.Context `inject:"context"`
}

func (inst *TestCommandHandler) _Impl() cli.Handler {
	return inst
}

func (inst *TestCommandHandler) Init(s cli.Service) error {
	return s.RegisterHandler(CommandArgsBypass, inst)
}

func (inst *TestCommandHandler) Handle(ctx *cli.TaskContext) error {
	console := ctx.Console
	cmd := ctx.CurrentTask.CommandName
	args := ctx.CurrentTask.Arguments
	console.WriteString("test command: " + cmd + "\n")
	for index, item := range args {
		console.WriteString(fmt.Sprint("args[", index, "] = ", item, "\n"))
	}

	bypass := GetBypassArgs(ctx.Context)
	return inst.check(ctx, bypass)
}

func (inst *TestCommandHandler) check(ctx *cli.TaskContext, bypass *BypassArgs) error {

	cmd1 := bypass.WantCmd
	args1 := bypass.WantArgs

	task := ctx.CurrentTask
	cmd2 := task.CommandName
	args2 := task.Arguments

	if cmd1 != cmd2 {
		want := ", want:" + cmd1
		have := ", have:" + cmd2
		return errors.New("bad command" + want + have)
	}

	a1 := inst.stringifyArgs(args1)
	a2 := inst.stringifyArgs(args2)
	if a1 != a2 {
		want := "\n want:" + a1
		have := "\n have:" + a2
		return errors.New("bad args" + want + have)
	}

	return nil
}

func (inst *TestCommandHandler) stringifyArgs(args []string) string {
	builder := strings.Builder{}
	for _, item := range args {
		builder.WriteString(" ")
		builder.WriteString(item)
	}
	return builder.String()
}

////////////////////////////////////////////////////////////////////////////////

type BypassArgs struct {
	WantCmd  string
	WantArgs []string
}

func GetBypassArgs(ctx context.Context) *BypassArgs {

	const name = "cli.test.BypassArgs"

	setter, err := contexts.GetContextSetter(ctx)
	if err != nil {
		panic(err)
	}

	o1 := setter.GetContext().Value(name)
	if o1 == nil {
		o2 := &BypassArgs{}
		setter.SetValue(name, o2)
		return o2
	}
	return o1.(*BypassArgs)
}

////////////////////////////////////////////////////////////////////////////////
