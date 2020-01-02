package usecase

// IAgent - interface for clients
type IAgent interface {
	GetIP() string
	GetOS() string
	GetVersion() string
}
