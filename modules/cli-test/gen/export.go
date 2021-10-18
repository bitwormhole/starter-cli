package gen

import "github.com/bitwormhole/starter/application"

func ExportConfigCliTest(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}
