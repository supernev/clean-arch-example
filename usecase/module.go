package usecase

// IModule - A game module interface
type IModule interface {
	Init()
	Run()
	Destroy()
}
