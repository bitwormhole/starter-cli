// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package test

import (
	cmds0x9b393b "github.com/bitwormhole/starter-cli/src/test/golang/cmds"
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComTestCommandHandler struct {
	instance *cmds0x9b393b.TestCommandHandler
	 markup0x23084a.Component `class:"cli-handler"`
	Context application0x67f6c5.Context `inject:"context"`
}

