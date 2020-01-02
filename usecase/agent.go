package usecase

// TypeOS - Types of OS
type TypeOS string

// OSType enum
const (
	OSApple   TypeOS = "Apple"
	OSAndroid TypeOS = "Android"
)

// IAgent - interface for clients
type IAgent interface {
	GetIP() string
	GetOS() TypeOS
	GetVersion() string

	OnRequestTrain(fn func(ballerID uint64, playerID uint64))
	ResponseTrain(ResultCode ErrorCode, OldLevel uint32, NewLevel uint32)
}
