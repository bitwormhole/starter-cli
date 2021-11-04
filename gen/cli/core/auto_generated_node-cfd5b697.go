// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package core

import (
	cli0xf30272 "github.com/bitwormhole/starter-cli/cli"
	support0xc85772 "github.com/bitwormhole/starter-cli/support"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComContext struct {
	instance *support0xc85772.Context
	 markup0x23084a.Component `id:"cli-context"`
	Service cli0xf30272.Service `inject:"#cli-service"`
	ClientFactory cli0xf30272.ClientFactory `inject:"#cli-client-factory"`
	Filters []cli0xf30272.Filter `inject:".cli-filter"`
	Handlers []cli0xf30272.Handler `inject:".cli-handler"`
}


type pComDefaultClientFactory struct {
	instance *support0xc85772.DefaultClientFactory
	 markup0x23084a.Component `id:"cli-client-factory"`
	CLI *support0xc85772.Context `inject:"#cli-context"`
}


type pComDefaultSerivce struct {
	instance *support0xc85772.DefaultSerivce
	 markup0x23084a.Component `id:"cli-service" initMethod:"Init"`
	CLI *support0xc85772.Context `inject:"#cli-context"`
	handlerTable map[string]cli0xf30272.Handler ``
	chain cli0xf30272.FilterChain ``
	chainBuilder *cli0xf30272.FilterChainBuilder ``
}

