package golang

import "testing"

func TestCLIContext(t *testing.T) {
	RunTest(t, func(tool *CLITesting) error {
		return tool.client.Execute("help", nil)
	})
}
