package terminal

import (
	"context"

	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/application"
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