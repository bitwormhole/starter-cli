package cli

import "strings"

// CommandLineBuilder 命令行创建器
type CommandLineBuilder struct {
	items []string
}

// AppendString 追加一段字符串
func (inst *CommandLineBuilder) AppendString(s string) {
	inst.items = append(inst.items, s)
}

// AppendStrings 追加好几段字符串
func (inst *CommandLineBuilder) AppendStrings(items []string) {
	for _, item := range items {
		inst.AppendString(item)
	}
}

// Create 创建命令行
func (inst *CommandLineBuilder) Create() string {

	items := inst.items
	if items == nil {
		return ""
	}

	builder := strings.Builder{}

	for _, item := range items {
		if builder.Len() > 0 {
			builder.WriteRune(' ')
		}
		if strings.ContainsAny(item, " \t") {
			mark := '"'
			if strings.ContainsRune(item, mark) {
				mark = '\''
			}
			builder.WriteRune(mark)
			builder.WriteString(item)
			builder.WriteRune(mark)
		} else {
			builder.WriteString(item)
		}
	}

	return builder.String()
}
