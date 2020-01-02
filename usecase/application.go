package usecase

// IApplication - Interface for a runnable application
type IApplication interface {
	AddModule(mod IModule)
	Start()
}
