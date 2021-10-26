package startercli

import (
	"github.com/bitwormhole/starter"
	gen1 "github.com/bitwormhole/starter-cli/modules/cli-core/gen"
	gen2 "github.com/bitwormhole/starter-cli/modules/cli-ext/gen"
	"github.com/bitwormhole/starter/application"
)

const (
	myModuleName = "github.com/bitwormhole/starter-cli"
	myModuleVer  = "v0.0.3"
	myModuleRev  = 3
)

// Module 导出模块【github.com/bitwormhole/starter-cli】
func Module() application.Module {

	mb := application.ModuleBuilder{}
	mb.Name(myModuleName).Version(myModuleVer).Revision(myModuleRev)
	mb.Resources(nil)
	mb.OnMount(gen1.ExportConfigCliCore)

	mb.Dependency(starter.Module())

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////

// ModuleWithBasicCommands 导出模块【github.com/bitwormhole/starter-cli#cmds】
func ModuleWithBasicCommands() application.Module {

	mb := application.ModuleBuilder{}
	mb.Name(myModuleName + "#cmds").Version(myModuleVer).Revision(myModuleRev)
	mb.Resources(nil)
	mb.OnMount(gen2.ExportConfigCliExt)

	mb.Dependency(Module())

	return mb.Create()
}
