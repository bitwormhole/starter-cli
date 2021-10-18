package handlers

import (
	"errors"

	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
)

type PWD struct {
	markup.Component `class:"cli-handler"`
}

func (inst *PWD) _Impl() cli.Handler {
	return inst
}

func (inst *PWD) Init(service cli.Service) error {
	return service.RegisterHandler("pwd", inst)
}

func (inst *PWD) Handle(t *cli.TaskContext) error {

	console := t.Console
	path := console.GetWorkingPath()
	path, err := normalizeWD(path)
	if err != nil {
		return err
	}

	console.WriteString(path.Path())
	console.WriteString("\n")

	return nil
}

func normalizeWD(path fs.Path) (fs.Path, error) {
	if path == nil {
		return nil, errors.New("path is nil")
	}
	p := path
	for ; p != nil; p = p.Parent() {
		if p.IsDir() {
			return p, nil
		}
	}
	return nil, errors.New("node is not exists, path=" + path.Path())
}

func (inst *PWD) GetHelpInfo() *cli.CommandHelpInfo {
	info := &cli.CommandHelpInfo{}
	info.Name = "pwd"
	info.Title = "显示工作目录"
	info.Description = "显示当前工作文件夹的路径"
	info.Content = "usage: pwd"
	return info
}
