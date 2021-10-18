package support

import (
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/markup"
)

type Context struct {
	markup.Component `id:"cli-context"`

	Service       cli.Service       `inject:"#cli-service"`
	ClientFactory cli.ClientFactory `inject:"#cli-client-factory"`
	Filters       []cli.Filter      `inject:".cli-filter"`
	Handlers      []cli.Handler     `inject:".cli-handler"`
}
