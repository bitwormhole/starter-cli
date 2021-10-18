package gen

import "github.com/bitwormhole/starter/application"

func ExportConfigCliExt(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}
