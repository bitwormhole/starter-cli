package handlers

import (
	"sort"
	"strconv"
	"strings"

	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/markup"
)

type Help struct {
	markup.Component `class:"cli-handler"`
}

func (inst *Help) _Impl() (cli.Handler, cli.CommandHelper) {
	return inst, inst
}

func (inst *Help) Init(service cli.Service) error {
	return service.RegisterHandler("help", inst)
}

func (inst *Help) Handle(t *cli.TaskContext) error {
	name, ok := inst.tryGetTargetCommandName(t)
	if ok {
		inst.displayDetailHelpString(t, name)
	} else {
		inst.displayOverviewHelpString(t)
	}
	return nil
}

func (inst *Help) tryGetTargetCommandName(t *cli.TaskContext) (string, bool) {
	task := t.CurrentTask
	args := task.Arguments
	if len(args) > 1 {
		return args[1], true
	}
	return "", false
}

func (inst *Help) displayOverviewHelpString(t *cli.TaskContext) {

	console := t.Console
	service := t.Service

	builder := strings.Builder{}
	namelist := service.GetHandlerNames()
	sort.Strings(namelist)

	for i, name := range namelist {
		inst.buildOverviewItem(i, name, service, &builder)
	}

	console.WriteString(builder.String())
}

func (inst *Help) buildOverviewItem(index int, name string, service cli.Service, sb *strings.Builder) {

	description := ""
	h, err := service.FindHandler(name)
	if err == nil {
		helper, ok := h.(cli.CommandHelper)
		if ok {
			info := helper.GetHelpInfo()
			description = info.Description
		}
	}

	sb.WriteString(strconv.Itoa(index))
	sb.WriteString(". ")
	sb.WriteString(inst.addPadding(name, 16))
	sb.WriteString(" ")
	sb.WriteString(description)

	sb.WriteString("\n")
}

func (inst *Help) addPadding(text string, width int) string {
	builder := strings.Builder{}
	builder.WriteString(text)
	for {
		if builder.Len() < width {
			builder.WriteRune(' ')
		} else {
			break
		}
	}
	return builder.String()
}

func (inst *Help) displayDetailHelpString(t *cli.TaskContext, cmd string) {

	console := t.Console
	service := t.Service

	h, err := service.FindHandler(cmd)
	if err != nil {
		console.WriteError("error: ", err)
		return
	}

	helper, ok := h.(cli.CommandHelper)
	if !ok {
		console.WriteString("no help info for command: " + cmd + "\n")
		return
	}

	info := helper.GetHelpInfo()
	console.WriteString(info.Name)
	console.WriteString(" (" + info.Title + ")\n")
	console.WriteString(info.Description + "\n")
	console.WriteString(info.Content + "\n")
}

func (inst *Help) GetHelpInfo() *cli.CommandHelpInfo {
	info := &cli.CommandHelpInfo{}
	info.Name = "help"
	info.Title = "帮助"
	info.Description = "显示帮助信息"
	info.Content = "usage: help [command_name]"
	return info
}
