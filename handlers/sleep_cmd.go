package handlers

import (
	"errors"
	"strconv"
	"time"

	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/markup"
)

type Sleep struct {
	markup.Component `class:"cli-handler"`
}

func (inst *Sleep) _Impl() cli.Handler {
	return inst
}

func (inst *Sleep) Init(service cli.Service) error {
	return service.RegisterHandler("sleep", inst)
}

func (inst *Sleep) Handle(t *cli.TaskContext) error {

	console := t.Console
	task := t.CurrentTask
	ms, err := inst.getArgMillisec(task.Arguments, 1)
	if err != nil {
		return err
	}

	console.WriteString("Sleep ")
	console.WriteString(strconv.Itoa(ms))
	console.WriteString(" ms ... ")
	time.Sleep(time.Millisecond * time.Duration(ms))
	console.WriteString("[done]\n")

	return nil
}

func (inst *Sleep) getArgMillisec(args []string, index int) (int, error) {
	if index < len(args) {
		text := args[index]
		return strconv.Atoi(text)
	}
	return 0, errors.New("bad argument")
}

func (inst *Sleep) GetHelpInfo() *cli.CommandHelpInfo {
	info := &cli.CommandHelpInfo{}
	info.Name = "sleep"
	info.Title = "睡一下"
	info.Description = "睡指定的毫秒"
	info.Content = "usage: sleep [ms]"
	return info
}
