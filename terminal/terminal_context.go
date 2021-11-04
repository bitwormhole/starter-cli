package terminal

import (
	"context"
	"os"

	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/contexts"
)

// Context 是终端（terminal）的上下文
type Context struct {
	client  cli.Client
	console cli.Console
	app     application.Context
	ctx     context.Context
	prompt  string
	exit    bool
}

// Prepare 准备命令行运行环境
func Prepare(ctx0 application.Context) (context.Context, error) {

	ctx := ctx0.NewChild()

	err := contexts.SetupApplicationContext(ctx)
	if err != nil {
		return nil, err
	}

	err = cli.SetupConsole(ctx, nil)
	if err != nil {
		return nil, err
	}

	console, err := cli.GetConsole(ctx)
	if err != nil {
		return nil, err
	}

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	console.SetOutput(os.Stdout)
	console.SetError(os.Stderr)
	console.SetInput(os.Stdin)
	console.SetWD(wd)

	return ctx, nil
}
