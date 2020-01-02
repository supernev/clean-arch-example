package usecase

import "time"

// ApplicationSimple - A simple implementation for a runnable application
type ApplicationSimple struct {
	modules []IModule
}

// NewApplicationSimple - Get new app instance
func NewApplicationSimple() *ApplicationSimple {
	return &ApplicationSimple{[]IModule{}}
}

// AddModule - Plugin a module
func (app *ApplicationSimple) AddModule(mod IModule) {
	app.modules = append(app.modules, mod)
}

// Start - main loop
func (app *ApplicationSimple) Start() {
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
