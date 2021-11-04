package golang

import (
	"testing"
)

func TestCmdHelp(t *testing.T) {
	tryTestCommand(t, "help")
}

func TestCmdHelp2(t *testing.T) {
	tryTestCommand(t, "help ls")
}

func TestCmdPWD(t *testing.T) {
	tryTestCommand(t, "pwd")
}

func TestCmdLS(t *testing.T) {
	tryTestCommand(t, "ls")
}

func TestCmdCD(t *testing.T) {
	tryTestCommand(t, "cd ./../../abc")
}

func TestCmdSleep(t *testing.T) {
	tryTestCommand(t, "sleep 5000")
}

////////////////////////////////////////////////////////////////////////////////

func tryTestCommand(t *testing.T, cmd string) {
	RunTest(t, func(tool *CLITesting) error {
		return tool.client.Execute(cmd, nil)
	})
}
