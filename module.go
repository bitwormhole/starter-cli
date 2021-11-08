package startercli

import (
	"embed"

	"github.com/bitwormhole/starter"
	gen1 "github.com/bitwormhole/starter-cli/gen/cli/core"
	gen2 "github.com/bitwormhole/starter-cli/gen/cli/ext"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	myModuleName = "github.com/bitwormhole/starter-cli"
	myModuleVer  = "v0.0.7"
	myModuleRev  = 7
)

//go:embed src/main/resources
var theMainRes embed.FS

// Module 导出模块【github.com/bitwormhole/starter-cli】
func Module() application.Module {

	mb := application.ModuleBuilder{}
	mb.Name(myModuleName).Version(myModuleVer).Revision(myModuleRev)
	mb.Resources(collection.LoadEmbedResources(&theMainRes, "src/main/resources"))
	mb.OnMount(gen1.ExportConfigCliCore)

	mb.Dependency(starter.Module())

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////

// ModuleWithBasicCommands 导出模块【github.com/bitwormhole/starter-cli#cmds】
func ModuleWithBasicCommands() application.Module {

	parent := Module()

	mb := application.ModuleBuilder{}
	mb.Name(parent.GetName() + "#cmds")
	mb.Version(parent.GetVersion())
	mb.Revision(parent.GetRevision())
	mb.Resources(parent.GetResources())
	mb.OnMount(gen2.ExportConfigCliExt)

	mb.Dependency(parent)

	return mb.Create()
}
