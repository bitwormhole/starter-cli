// (todo:gen2.template)
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package ext

import (
	handlers0xed7f6d "github.com/bitwormhole/starter-cli/handlers"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
)

func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()

	// component: com0-handlers0xed7f6d.CD
	cominfobuilder.Next()
	cominfobuilder.ID("com0-handlers0xed7f6d.CD").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComCD{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com1-handlers0xed7f6d.Help
	cominfobuilder.Next()
	cominfobuilder.ID("com1-handlers0xed7f6d.Help").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComHelp{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com2-handlers0xed7f6d.LS
	cominfobuilder.Next()
	cominfobuilder.ID("com2-handlers0xed7f6d.LS").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLS{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com3-handlers0xed7f6d.PWD
	cominfobuilder.Next()
	cominfobuilder.ID("com3-handlers0xed7f6d.PWD").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPWD{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com4-handlers0xed7f6d.Sleep
	cominfobuilder.Next()
	cominfobuilder.ID("com4-handlers0xed7f6d.Sleep").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComSleep{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComCD : the factory of component: com0-handlers0xed7f6d.CD
type comFactory4pComCD struct {
	mPrototype *handlers0xed7f6d.CD
}

func (inst *comFactory4pComCD) init() application.ComponentFactory {

	inst.mPrototype = inst.newObject()
	return inst
}

func (inst *comFactory4pComCD) newObject() *handlers0xed7f6d.CD {
	return &handlers0xed7f6d.CD{}
}

func (inst *comFactory4pComCD) castObject(instance application.ComponentInstance) *handlers0xed7f6d.CD {
	return instance.Get().(*handlers0xed7f6d.CD)
}

func (inst *comFactory4pComCD) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst *comFactory4pComCD) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst *comFactory4pComCD) AfterService() application.ComponentAfterService {
	return inst
}

func (inst *comFactory4pComCD) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComCD) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComCD) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComHelp : the factory of component: com1-handlers0xed7f6d.Help
type comFactory4pComHelp struct {
	mPrototype *handlers0xed7f6d.Help
}

func (inst *comFactory4pComHelp) init() application.ComponentFactory {

	inst.mPrototype = inst.newObject()
	return inst
}

func (inst *comFactory4pComHelp) newObject() *handlers0xed7f6d.Help {
	return &handlers0xed7f6d.Help{}
}

func (inst *comFactory4pComHelp) castObject(instance application.ComponentInstance) *handlers0xed7f6d.Help {
	return instance.Get().(*handlers0xed7f6d.Help)
}

func (inst *comFactory4pComHelp) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst *comFactory4pComHelp) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst *comFactory4pComHelp) AfterService() application.ComponentAfterService {
	return inst
}

func (inst *comFactory4pComHelp) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComHelp) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComHelp) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComLS : the factory of component: com2-handlers0xed7f6d.LS
type comFactory4pComLS struct {
	mPrototype *handlers0xed7f6d.LS
}

func (inst *comFactory4pComLS) init() application.ComponentFactory {

	inst.mPrototype = inst.newObject()
	return inst
}

func (inst *comFactory4pComLS) newObject() *handlers0xed7f6d.LS {
	return &handlers0xed7f6d.LS{}
}

func (inst *comFactory4pComLS) castObject(instance application.ComponentInstance) *handlers0xed7f6d.LS {
	return instance.Get().(*handlers0xed7f6d.LS)
}

func (inst *comFactory4pComLS) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst *comFactory4pComLS) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst *comFactory4pComLS) AfterService() application.ComponentAfterService {
	return inst
}

func (inst *comFactory4pComLS) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComLS) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComLS) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComPWD : the factory of component: com3-handlers0xed7f6d.PWD
type comFactory4pComPWD struct {
	mPrototype *handlers0xed7f6d.PWD
}

func (inst *comFactory4pComPWD) init() application.ComponentFactory {

	inst.mPrototype = inst.newObject()
	return inst
}

func (inst *comFactory4pComPWD) newObject() *handlers0xed7f6d.PWD {
	return &handlers0xed7f6d.PWD{}
}

func (inst *comFactory4pComPWD) castObject(instance application.ComponentInstance) *handlers0xed7f6d.PWD {
	return instance.Get().(*handlers0xed7f6d.PWD)
}

func (inst *comFactory4pComPWD) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst *comFactory4pComPWD) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst *comFactory4pComPWD) AfterService() application.ComponentAfterService {
	return inst
}

func (inst *comFactory4pComPWD) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComPWD) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComPWD) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComSleep : the factory of component: com4-handlers0xed7f6d.Sleep
type comFactory4pComSleep struct {
	mPrototype *handlers0xed7f6d.Sleep
}

func (inst *comFactory4pComSleep) init() application.ComponentFactory {

	inst.mPrototype = inst.newObject()
	return inst
}

func (inst *comFactory4pComSleep) newObject() *handlers0xed7f6d.Sleep {
	return &handlers0xed7f6d.Sleep{}
}

func (inst *comFactory4pComSleep) castObject(instance application.ComponentInstance) *handlers0xed7f6d.Sleep {
	return instance.Get().(*handlers0xed7f6d.Sleep)
}

func (inst *comFactory4pComSleep) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst *comFactory4pComSleep) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst *comFactory4pComSleep) AfterService() application.ComponentAfterService {
	return inst
}

func (inst *comFactory4pComSleep) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComSleep) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComSleep) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}
