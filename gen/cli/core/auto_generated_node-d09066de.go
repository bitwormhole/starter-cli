// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package core

import (
	context0xec2727 "context"
	cli0xf30272 "github.com/bitwormhole/starter-cli/cli"
	filters0x54a607 "github.com/bitwormhole/starter-cli/filters"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComHandlerFinderFilter struct {
	instance *filters0x54a607.HandlerFinderFilter
	 markup0x23084a.Component `class:"cli-filter"`
	Priority int `inject:"800"`
}


type pComContextFilter struct {
	instance *filters0x54a607.ContextFilter
	 markup0x23084a.Component `class:"cli-filter"`
	Priority int `inject:"900"`
	Context context0xec2727.Context `inject:"context"`
	Service cli0xf30272.Service ``
}


type pComExecutorFilter struct {
	instance *filters0x54a607.ExecutorFilter
	 markup0x23084a.Component `class:"cli-filter"`
	Priority int `inject:"700"`
}


type pComMultilineSupportFilter struct {
	instance *filters0x54a607.MultilineSupportFilter
	 markup0x23084a.Component `class:"cli-filter"`
	Priority int `inject:"850"`
}


type pComNopFilter struct {
	instance *filters0x54a607.NopFilter
	 markup0x23084a.Component `class:"cli-filter"`
	Priority int `inject:"0"`
}

