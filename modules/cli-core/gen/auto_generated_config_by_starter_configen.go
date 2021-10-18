// (todo:gen2.template)
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gen

import (
	cli0xf30272 "github.com/bitwormhole/starter-cli/cli"
	filters0x54a607 "github.com/bitwormhole/starter-cli/filters"
	support0xc85772 "github.com/bitwormhole/starter-cli/support"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
)

func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()

	// component: cli-context
	cominfobuilder.Next()
	cominfobuilder.ID("cli-context").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComContext{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: cli-client-factory
	cominfobuilder.Next()
	cominfobuilder.ID("cli-client-factory").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDefaultClientFactory{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: cli-service
	cominfobuilder.Next()
	cominfobuilder.ID("cli-service").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDefaultSerivce{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com3-filters0x54a607.HandlerFinderFilter
	cominfobuilder.Next()
	cominfobuilder.ID("com3-filters0x54a607.HandlerFinderFilter").Class("cli-filter").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComHandlerFinderFilter{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com4-filters0x54a607.ContextFilter
	cominfobuilder.Next()
	cominfobuilder.ID("com4-filters0x54a607.ContextFilter").Class("cli-filter").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComContextFilter{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com5-filters0x54a607.ExecutorFilter
	cominfobuilder.Next()
	cominfobuilder.ID("com5-filters0x54a607.ExecutorFilter").Class("cli-filter").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExecutorFilter{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com6-filters0x54a607.MultilineSupportFilter
	cominfobuilder.Next()
	cominfobuilder.ID("com6-filters0x54a607.MultilineSupportFilter").Class("cli-filter").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComMultilineSupportFilter{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com7-filters0x54a607.NopFilter
	cominfobuilder.Next()
	cominfobuilder.ID("com7-filters0x54a607.NopFilter").Class("cli-filter").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComNopFilter{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComContext : the factory of component: cli-context
type comFactory4pComContext struct {
	mPrototype *support0xc85772.Context

	mServiceSelector       config.InjectionSelector
	mClientFactorySelector config.InjectionSelector
	mFiltersSelector       config.InjectionSelector
	mHandlersSelector      config.InjectionSelector
}

func (inst *comFactory4pComContext) init() application.ComponentFactory {

	inst.mServiceSelector = config.NewInjectionSelector("#cli-service", nil)
	inst.mClientFactorySelector = config.NewInjectionSelector("#cli-client-factory", nil)
	inst.mFiltersSelector = config.NewInjectionSelector(".cli-filter", nil)
	inst.mHandlersSelector = config.NewInjectionSelector(".cli-handler", nil)

	inst.mPrototype = inst.newObject()
	return inst
}

func (inst *comFactory4pComContext) newObject() *support0xc85772.Context {
	return &support0xc85772.Context{}
}

func (inst *comFactory4pComContext) castObject(instance application.ComponentInstance) *support0xc85772.Context {
	return instance.Get().(*support0xc85772.Context)
}

func (inst *comFactory4pComContext) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst *comFactory4pComContext) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst *comFactory4pComContext) AfterService() application.ComponentAfterService {
	return inst
}

func (inst *comFactory4pComContext) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComContext) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComContext) Inject(instance application.ComponentInstance, context application.InstanceContext) error {

	obj := inst.castObject(instance)
	obj.Service = inst.getterForFieldServiceSelector(context)
	obj.ClientFactory = inst.getterForFieldClientFactorySelector(context)
	obj.Filters = inst.getterForFieldFiltersSelector(context)
	obj.Handlers = inst.getterForFieldHandlersSelector(context)
	return context.LastError()
}

//getterForFieldServiceSelector
func (inst *comFactory4pComContext) getterForFieldServiceSelector(context application.InstanceContext) cli0xf30272.Service {

	o1 := inst.mServiceSelector.GetOne(context)
	o2, ok := o1.(cli0xf30272.Service)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "cli-context")
		eb.Set("field", "Service")
		eb.Set("type1", "?")
		eb.Set("type2", "cli0xf30272.Service")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldClientFactorySelector
func (inst *comFactory4pComContext) getterForFieldClientFactorySelector(context application.InstanceContext) cli0xf30272.ClientFactory {

	o1 := inst.mClientFactorySelector.GetOne(context)
	o2, ok := o1.(cli0xf30272.ClientFactory)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "cli-context")
		eb.Set("field", "ClientFactory")
		eb.Set("type1", "?")
		eb.Set("type2", "cli0xf30272.ClientFactory")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldFiltersSelector
func (inst *comFactory4pComContext) getterForFieldFiltersSelector(context application.InstanceContext) []cli0xf30272.Filter {
	list1 := inst.mFiltersSelector.GetList(context)
	list2 := make([]cli0xf30272.Filter, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(cli0xf30272.Filter)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}

//getterForFieldHandlersSelector
func (inst *comFactory4pComContext) getterForFieldHandlersSelector(context application.InstanceContext) []cli0xf30272.Handler {
	list1 := inst.mHandlersSelector.GetList(context)
	list2 := make([]cli0xf30272.Handler, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(cli0xf30272.Handler)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDefaultClientFactory : the factory of component: cli-client-factory
type comFactory4pComDefaultClientFactory struct {
	mPrototype *support0xc85772.DefaultClientFactory

	mCLISelector config.InjectionSelector
}

func (inst *comFactory4pComDefaultClientFactory) init() application.ComponentFactory {

	inst.mCLISelector = config.NewInjectionSelector("#cli-context", nil)

	inst.mPrototype = inst.newObject()
	return inst
}

func (inst *comFactory4pComDefaultClientFactory) newObject() *support0xc85772.DefaultClientFactory {
	return &support0xc85772.DefaultClientFactory{}
}

func (inst *comFactory4pComDefaultClientFactory) castObject(instance application.ComponentInstance) *support0xc85772.DefaultClientFactory {
	return instance.Get().(*support0xc85772.DefaultClientFactory)
}

func (inst *comFactory4pComDefaultClientFactory) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst *comFactory4pComDefaultClientFactory) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst *comFactory4pComDefaultClientFactory) AfterService() application.ComponentAfterService {
	return inst
}

func (inst *comFactory4pComDefaultClientFactory) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComDefaultClientFactory) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComDefaultClientFactory) Inject(instance application.ComponentInstance, context application.InstanceContext) error {

	obj := inst.castObject(instance)
	obj.CLI = inst.getterForFieldCLISelector(context)
	return context.LastError()
}

//getterForFieldCLISelector
func (inst *comFactory4pComDefaultClientFactory) getterForFieldCLISelector(context application.InstanceContext) *support0xc85772.Context {

	o1 := inst.mCLISelector.GetOne(context)
	o2, ok := o1.(*support0xc85772.Context)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "cli-client-factory")
		eb.Set("field", "CLI")
		eb.Set("type1", "?")
		eb.Set("type2", "*support0xc85772.Context")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDefaultSerivce : the factory of component: cli-service
type comFactory4pComDefaultSerivce struct {
	mPrototype *support0xc85772.DefaultSerivce

	mCLISelector config.InjectionSelector
}

func (inst *comFactory4pComDefaultSerivce) init() application.ComponentFactory {

	inst.mCLISelector = config.NewInjectionSelector("#cli-context", nil)

	inst.mPrototype = inst.newObject()
	return inst
}

func (inst *comFactory4pComDefaultSerivce) newObject() *support0xc85772.DefaultSerivce {
	return &support0xc85772.DefaultSerivce{}
}

func (inst *comFactory4pComDefaultSerivce) castObject(instance application.ComponentInstance) *support0xc85772.DefaultSerivce {
	return instance.Get().(*support0xc85772.DefaultSerivce)
}

func (inst *comFactory4pComDefaultSerivce) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst *comFactory4pComDefaultSerivce) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst *comFactory4pComDefaultSerivce) AfterService() application.ComponentAfterService {
	return inst
}

func (inst *comFactory4pComDefaultSerivce) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst *comFactory4pComDefaultSerivce) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComDefaultSerivce) Inject(instance application.ComponentInstance, context application.InstanceContext) error {

	obj := inst.castObject(instance)
	obj.CLI = inst.getterForFieldCLISelector(context)
	return context.LastError()
}

//getterForFieldCLISelector
func (inst *comFactory4pComDefaultSerivce) getterForFieldCLISelector(context application.InstanceContext) *support0xc85772.Context {

	o1 := inst.mCLISelector.GetOne(context)
	o2, ok := o1.(*support0xc85772.Context)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "cli-service")
		eb.Set("field", "CLI")
		eb.Set("type1", "?")
		eb.Set("type2", "*support0xc85772.Context")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComHandlerFinderFilter : the factory of component: com3-filters0x54a607.HandlerFinderFilter
type comFactory4pComHandlerFinderFilter struct {
	mPrototype *filters0x54a607.HandlerFinderFilter

	mPrioritySelector config.InjectionSelector
}

func (inst *comFactory4pComHandlerFinderFilter) init() application.ComponentFactory {

	inst.mPrioritySelector = config.NewInjectionSelector("800", nil)

	inst.mPrototype = inst.newObject()
	return inst
}

func (inst *comFactory4pComHandlerFinderFilter) newObject() *filters0x54a607.HandlerFinderFilter {
	return &filters0x54a607.HandlerFinderFilter{}
}

func (inst *comFactory4pComHandlerFinderFilter) castObject(instance application.ComponentInstance) *filters0x54a607.HandlerFinderFilter {
	return instance.Get().(*filters0x54a607.HandlerFinderFilter)
}

func (inst *comFactory4pComHandlerFinderFilter) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst *comFactory4pComHandlerFinderFilter) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst *comFactory4pComHandlerFinderFilter) AfterService() application.ComponentAfterService {
	return inst
}

func (inst *comFactory4pComHandlerFinderFilter) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComHandlerFinderFilter) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComHandlerFinderFilter) Inject(instance application.ComponentInstance, context application.InstanceContext) error {

	obj := inst.castObject(instance)
	obj.Priority = inst.getterForFieldPrioritySelector(context)
	return context.LastError()
}

//getterForFieldPrioritySelector
func (inst *comFactory4pComHandlerFinderFilter) getterForFieldPrioritySelector(context application.InstanceContext) int {
	return inst.mPrioritySelector.GetInt(context)
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComContextFilter : the factory of component: com4-filters0x54a607.ContextFilter
type comFactory4pComContextFilter struct {
	mPrototype *filters0x54a607.ContextFilter

	mPrioritySelector config.InjectionSelector
	mContextSelector  config.InjectionSelector
}

func (inst *comFactory4pComContextFilter) init() application.ComponentFactory {

	inst.mPrioritySelector = config.NewInjectionSelector("900", nil)
	inst.mContextSelector = config.NewInjectionSelector("context", nil)

	inst.mPrototype = inst.newObject()
	return inst
}

func (inst *comFactory4pComContextFilter) newObject() *filters0x54a607.ContextFilter {
	return &filters0x54a607.ContextFilter{}
}

func (inst *comFactory4pComContextFilter) castObject(instance application.ComponentInstance) *filters0x54a607.ContextFilter {
	return instance.Get().(*filters0x54a607.ContextFilter)
}

func (inst *comFactory4pComContextFilter) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst *comFactory4pComContextFilter) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst *comFactory4pComContextFilter) AfterService() application.ComponentAfterService {
	return inst
}

func (inst *comFactory4pComContextFilter) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComContextFilter) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComContextFilter) Inject(instance application.ComponentInstance, context application.InstanceContext) error {

	obj := inst.castObject(instance)
	obj.Priority = inst.getterForFieldPrioritySelector(context)
	obj.Context = inst.getterForFieldContextSelector(context)
	return context.LastError()
}

//getterForFieldPrioritySelector
func (inst *comFactory4pComContextFilter) getterForFieldPrioritySelector(context application.InstanceContext) int {
	return inst.mPrioritySelector.GetInt(context)
}

//getterForFieldContextSelector
func (inst *comFactory4pComContextFilter) getterForFieldContextSelector(context application.InstanceContext) application.Context {
	return context.Context()
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExecutorFilter : the factory of component: com5-filters0x54a607.ExecutorFilter
type comFactory4pComExecutorFilter struct {
	mPrototype *filters0x54a607.ExecutorFilter

	mPrioritySelector config.InjectionSelector
}

func (inst *comFactory4pComExecutorFilter) init() application.ComponentFactory {

	inst.mPrioritySelector = config.NewInjectionSelector("700", nil)

	inst.mPrototype = inst.newObject()
	return inst
}

func (inst *comFactory4pComExecutorFilter) newObject() *filters0x54a607.ExecutorFilter {
	return &filters0x54a607.ExecutorFilter{}
}

func (inst *comFactory4pComExecutorFilter) castObject(instance application.ComponentInstance) *filters0x54a607.ExecutorFilter {
	return instance.Get().(*filters0x54a607.ExecutorFilter)
}

func (inst *comFactory4pComExecutorFilter) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst *comFactory4pComExecutorFilter) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst *comFactory4pComExecutorFilter) AfterService() application.ComponentAfterService {
	return inst
}

func (inst *comFactory4pComExecutorFilter) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComExecutorFilter) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComExecutorFilter) Inject(instance application.ComponentInstance, context application.InstanceContext) error {

	obj := inst.castObject(instance)
	obj.Priority = inst.getterForFieldPrioritySelector(context)
	return context.LastError()
}

//getterForFieldPrioritySelector
func (inst *comFactory4pComExecutorFilter) getterForFieldPrioritySelector(context application.InstanceContext) int {
	return inst.mPrioritySelector.GetInt(context)
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComMultilineSupportFilter : the factory of component: com6-filters0x54a607.MultilineSupportFilter
type comFactory4pComMultilineSupportFilter struct {
	mPrototype *filters0x54a607.MultilineSupportFilter

	mPrioritySelector config.InjectionSelector
}

func (inst *comFactory4pComMultilineSupportFilter) init() application.ComponentFactory {

	inst.mPrioritySelector = config.NewInjectionSelector("850", nil)

	inst.mPrototype = inst.newObject()
	return inst
}

func (inst *comFactory4pComMultilineSupportFilter) newObject() *filters0x54a607.MultilineSupportFilter {
	return &filters0x54a607.MultilineSupportFilter{}
}

func (inst *comFactory4pComMultilineSupportFilter) castObject(instance application.ComponentInstance) *filters0x54a607.MultilineSupportFilter {
	return instance.Get().(*filters0x54a607.MultilineSupportFilter)
}

func (inst *comFactory4pComMultilineSupportFilter) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst *comFactory4pComMultilineSupportFilter) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst *comFactory4pComMultilineSupportFilter) AfterService() application.ComponentAfterService {
	return inst
}

func (inst *comFactory4pComMultilineSupportFilter) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComMultilineSupportFilter) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComMultilineSupportFilter) Inject(instance application.ComponentInstance, context application.InstanceContext) error {

	obj := inst.castObject(instance)
	obj.Priority = inst.getterForFieldPrioritySelector(context)
	return context.LastError()
}

//getterForFieldPrioritySelector
func (inst *comFactory4pComMultilineSupportFilter) getterForFieldPrioritySelector(context application.InstanceContext) int {
	return inst.mPrioritySelector.GetInt(context)
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComNopFilter : the factory of component: com7-filters0x54a607.NopFilter
type comFactory4pComNopFilter struct {
	mPrototype *filters0x54a607.NopFilter

	mPrioritySelector config.InjectionSelector
}

func (inst *comFactory4pComNopFilter) init() application.ComponentFactory {

	inst.mPrioritySelector = config.NewInjectionSelector("0", nil)

	inst.mPrototype = inst.newObject()
	return inst
}

func (inst *comFactory4pComNopFilter) newObject() *filters0x54a607.NopFilter {
	return &filters0x54a607.NopFilter{}
}

func (inst *comFactory4pComNopFilter) castObject(instance application.ComponentInstance) *filters0x54a607.NopFilter {
	return instance.Get().(*filters0x54a607.NopFilter)
}

func (inst *comFactory4pComNopFilter) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst *comFactory4pComNopFilter) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst *comFactory4pComNopFilter) AfterService() application.ComponentAfterService {
	return inst
}

func (inst *comFactory4pComNopFilter) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComNopFilter) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComNopFilter) Inject(instance application.ComponentInstance, context application.InstanceContext) error {

	obj := inst.castObject(instance)
	obj.Priority = inst.getterForFieldPrioritySelector(context)
	return context.LastError()
}

//getterForFieldPrioritySelector
func (inst *comFactory4pComNopFilter) getterForFieldPrioritySelector(context application.InstanceContext) int {
	return inst.mPrioritySelector.GetInt(context)
}
