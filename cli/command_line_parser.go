package cli

import (
	"errors"
	"strings"
)

// CommandLineParser 命令行解析器
type CommandLineParser struct {
}

// Parse 解析命令行
func (inst *CommandLineParser) Parse(line string) ([]string, error) {

	args := make([]string, 0)
	reader := &commandLineReader{}
	reader.init(line)

	for {
		reader.skipSpace()
		if reader.hasMore() {
			element, err := reader.read()
			if err != nil {
				return nil, err
			}
			if element != "" {
				args = append(args, element)
			}
		} else {
			break
		}
	}

	return args, nil
}

////////////////////////////////////////////////////////////////////////////////

type commandLineReader struct {
	source  []rune
	ptr     int
	length  int
	builder strings.Builder
}

func (inst *commandLineReader) init(line string) {
	array := ([]rune(line))
	inst.ptr = 0
	inst.length = len(array)
	inst.source = array
}

func (inst *commandLineReader) hasMore() bool {
	return inst.ptr < inst.length
}

func (inst *commandLineReader) read() (string, error) {

	i := inst.ptr
	length := inst.length
	array := inst.source

	if i >= length {
		return "", errors.New("EOF")
	}

	const mk1 = '\''
	const mk2 = '"'
	ch := array[i]

	if ch == '=' {
		inst.ptr = i + 1
		return "", nil
	} else if ch == mk1 {
		return inst.readText(mk1)
	} else if ch == mk2 {
		return inst.readText(mk2)
	} else {
		return inst.readToken()
	}
}

func (inst *commandLineReader) readToken() (string, error) {

	builder := &inst.builder
	i := inst.ptr
	length := inst.length
	array := inst.source
	chs := []rune{' ', '\t', '=', '"', '\'', '\n', '\r'} // ending of token

	builder.Reset()

	for ; i < length; i++ {
		ch := array[i]
		if inst.containsRune(ch, chs) {
			break
		}
		builder.WriteRune(ch)
	}

	inst.ptr = i
	return builder.String(), nil
}

func (inst *commandLineReader) readText(ending rune) (string, error) {

	builder := &inst.builder
	i := inst.ptr + 1
	length := inst.length
	array := inst.source

	builder.Reset()

	for ; i < length; i++ {
		ch := array[i]
		if ch == ending {
			inst.ptr = i + 1
			return builder.String(), nil
		}
		builder.WriteRune(ch)
	}

	return "", errors.New("end with exception")
}

func (inst *commandLineReader) skipSpace() {
	i := inst.ptr
	length := inst.length
	array := inst.source
	chs := []rune{' ', '\t', '\n', '\r'}
	for ; i < length; i++ {
		ch := array[i]
		if inst.containsRune(ch, chs) {
			continue
		} else {
			break
		}
	}
	inst.ptr = i
}

func (inst *commandLineReader) containsRune(ch rune, list []rune) bool {
	if list == nil {
		return false
	}
	for _, ch2 := range list {
		if ch == ch2 {
			return true
		}
	}
	return false
}
