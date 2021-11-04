// (todo:gen2.template)
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package test

import (
	cmds0x9b393b "github.com/bitwormhole/starter-cli/src/test/golang/cmds"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
)

func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()

	// component: com0-cmds0x9b393b.TestCommandHandler
	cominfobuilder.Next()
	cominfobuilder.ID("com0-cmds0x9b393b.TestCommandHandler").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComTestCommandHandler{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComTestCommandHandler : the factory of component: com0-cmds0x9b393b.TestCommandHandler
type comFactory4pComTestCommandHandler struct {
	mPrototype *cmds0x9b393b.TestCommandHandler

	mContextSelector config.InjectionSelector
}

func (inst *comFactory4pComTestCommandHandler) init() application.ComponentFactory {

	inst.mContextSelector = config.NewInjectionSelector("context", nil)

	inst.mPrototype = inst.newObject()
	return inst
}

func (inst *comFactory4pComTestCommandHandler) newObject() *cmds0x9b393b.TestCommandHandler {
	return &cmds0x9b393b.TestCommandHandler{}
}

func (inst *comFactory4pComTestCommandHandler) castObject(instance application.ComponentInstance) *cmds0x9b393b.TestCommandHandler {
	return instance.Get().(*cmds0x9b393b.TestCommandHandler)
}

func (inst *comFactory4pComTestCommandHandler) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst *comFactory4pComTestCommandHandler) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst *comFactory4pComTestCommandHandler) AfterService() application.ComponentAfterService {
	return inst
}

func (inst *comFactory4pComTestCommandHandler) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComTestCommandHandler) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComTestCommandHandler) Inject(instance application.ComponentInstance, context application.InstanceContext) error {

	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst *comFactory4pComTestCommandHandler) getterForFieldContextSelector(context application.InstanceContext) application.Context {
	return context.Context()
}
