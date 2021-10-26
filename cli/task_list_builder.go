package cli

import (
	"strings"
)

// TaskListBuilder ...
type TaskListBuilder struct {
	list []*TaskUnit
}

func (inst *TaskListBuilder) getList() []*TaskUnit {
	list := inst.list
	if list == nil {
		list = make([]*TaskUnit, 0)
		inst.list = list
	}
	return list
}

// Create ...
func (inst *TaskListBuilder) Create() []*TaskUnit {
	src := inst.getList()
	dst := make([]*TaskUnit, 0)
	for index, item := range src {
		item.LineNumber = index + 1
		args := item.Arguments
		if args == nil {
			continue
		}
		if len(args) == 0 {
			continue
		}
		dst = append(dst, item)
	}
	return dst
}

// AddLine @line: 完整的行文本； @index: 行号（base=0）； @args: 命令以及参数；
func (inst *TaskListBuilder) AddLine(line string, index int, args []string) error {

	if line == "" {
		clb := CommandLineBuilder{}
		clb.AppendStrings(args)
		line = clb.Create()
	}

	unit := &TaskUnit{}
	unit.LineNumber = index + 1
	unit.Arguments = args
	unit.CommandLine = line

	list := inst.getList()
	list = append(list, unit)
	inst.list = list
	return nil
}

func (inst *TaskListBuilder) parseLine(line string, index int) error {
	parser := CommandLineParser{}
	args, err := parser.Parse(line)
	if err != nil {
		return err
	}
	return inst.AddLine(line, index, args)
}

// ParseScript 解析脚本
func (inst *TaskListBuilder) ParseScript(script string) error {
	const ch1 = "\r"
	const ch2 = "\n"
	script = strings.ReplaceAll(script, ch1, ch2)
	array := strings.Split(script, ch2)
	for index, line := range array {
		err := inst.parseLine(line, index)
		if err != nil {
			return err
		}
	}
	return nil
}
