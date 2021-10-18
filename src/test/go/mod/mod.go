package mod

import (
	startercli "github.com/bitwormhole/starter-cli"
	"github.com/bitwormhole/starter/application"
)

// ExportModuleForCLITest 导出测试模块
func ExportModuleForCLITest() application.Module {

	parent := startercli.Module()

	mb := application.ModuleBuilder{}
	mb.Name(parent.GetName() + "#test")
	mb.Version(parent.GetVersion())
	mb.Revision(parent.GetRevision())

	mb.Dependency(parent)
	mb.Dependency(startercli.ModuleWithBasicCommands())

	return mb.Create()
}
