package usecase

// SimpleEventContext - An implementation of IEventContext
type SimpleEventContext struct {
	params map[string]interface{}
}

// NewSimpleEventContext - create a simple context instance
func NewSimpleEventContext() *SimpleEventContext {
	return &SimpleEventContext{
		params: map[string]interface{}{},
	}
}

// Get - get a param of any type by name
func (ctx *SimpleEventContext) Get(name string) interface{} {
	return ctx.params[name]
}

// Set - set a param of any type by name
func (ctx *SimpleEventContext) Set(name string, value interface{}) {
	ctx.params[name] = value
}
