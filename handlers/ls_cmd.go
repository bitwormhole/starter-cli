package handlers

import (
	"errors"
	"strconv"
	"strings"

	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
)

type LS struct {
	markup.Component `class:"cli-handler"`
}

func (inst *LS) _Impl() cli.Handler {
	return inst
}

func (inst *LS) Init(service cli.Service) error {
	return service.RegisterHandler("ls", inst)
}

func (inst *LS) Handle(t *cli.TaskContext) error {

	console := t.Console
	path := console.GetWorkingPath()
	path, err := normalizeWD(path)
	if err != nil {
		return err
	}

	if !path.IsDir() {
		return errors.New("it is not a directory, path=" + path.Path())
	}

	list := path.ListItems()
	for _, item := range list {
		str := inst.stringifyItem(item)
		console.WriteString(str)
		console.WriteString("\n")
	}

	return nil
}

func (inst *LS) stringifyItem(item fs.Path) string {

	builder := strings.Builder{}

	meta := item.GetMeta()
	size := meta.Size()
	date := meta.LastModTime()
	name := item.Name()
	mode := meta.Mode()

	var strSize, strTime string
	strTime = inst.stringifyTime(date)

	if meta.IsDir() {
		strSize = ("[dir]")
		strSize = inst.addPaddingToString(strSize, 10, false)
	} else {
		strSize = strconv.FormatInt(size, 10)
		strSize = inst.addPaddingToString(strSize, 10, true)
	}

	strTime = inst.addPaddingToString(strTime, 10, false)

	builder.WriteString(mode.String())
	builder.WriteString(" ")
	builder.WriteString(strTime)
	builder.WriteString(" ")
	builder.WriteString(strSize)
	builder.WriteString(" ")
	builder.WriteString(name)

	return builder.String()
}

func (inst *LS) addPaddingToString(str string, width int, alignRight bool) string {
	len := len(str)
	const space = " "
	for ; len < width; len++ {
		if alignRight {
			str = space + str
		} else {
			str = str + space
		}
	}
	return str
}

func (inst *LS) stringifyTime(t int64) string {
	t2 := util.Int64ToTime(t)
	return t2.String()
}

func (inst *LS) GetHelpInfo() *cli.CommandHelpInfo {
	info := &cli.CommandHelpInfo{}
	info.Name = "ls"
	info.Title = "显示文件夹内容"
	info.Description = "显示文件夹内容"
	info.Content = "usage: ls [模式]"
	return info
}
