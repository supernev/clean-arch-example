package usecase

import "time"

// SimpleApplication - A simple implementation for a runnable application
type SimpleApplication struct {
	modules []IModule
}

// NewSimpleApplication - Get new app instance
func NewSimpleApplication() *SimpleApplication {
	return &SimpleApplication{[]IModule{}}
}

// AddModule - Plugin a module
func (app *SimpleApplication) AddModule(mod IModule) {
	app.modules = append(app.modules, mod)
}

// Start - main loop
func (app *SimpleApplication) Start() {
	for _, mod := range app.modules {
		mod.Init()
		defer mod.Destroy()
	}

	for {
		for _, mod := range app.modules {
			mod.Run()
		}
		time.Sleep(1000)
	}
}
