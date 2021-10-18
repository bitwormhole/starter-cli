package gen

import "github.com/bitwormhole/starter/application"

func ExportConfigCliCore(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}
