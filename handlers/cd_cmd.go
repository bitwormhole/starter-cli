package handlers

import (
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
)

type CD struct {
	markup.Component `class:"cli-handler"`
}

func (inst *CD) _Impl() cli.Handler {
	return inst
}

func (inst *CD) Init(service cli.Service) error {
	return service.RegisterHandler("cd", inst)
}

func (inst *CD) Handle(t *cli.TaskContext) error {

	console := t.Console
	path := console.GetWorkingPath()
	task := t.CurrentTask

	to, err := inst.getArgTo(task.Arguments)
	if err != nil {
		return err
	}

	if inst.isAbsolutePath(to, path.FileSystem()) {
		path = path.FileSystem().GetPath(to)
	} else {
		path = path.GetChild(to)
	}

	if path.IsDir() {
		console.SetWorkingPath(path)
	} else {
		console.WriteError("The path is not a directory: "+path.Path(), nil)
	}

	return nil
}

func (inst *CD) getArgTo(args []string) (string, error) {
	return args[1], nil
}

func (inst *CD) isAbsolutePath(path string, fs fs.FileSystem) bool {
	return fs.IsAbsolute(path)
}

func (inst *CD) GetHelpInfo() *cli.CommandHelpInfo {
	info := &cli.CommandHelpInfo{}
	info.Name = "cd"
	info.Title = "切换工作目录"
	info.Description = "切换当前工作文件夹的路径"
	info.Content = "usage: cd [target]"
	return info
}
