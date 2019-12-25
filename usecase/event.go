package usecase

// EventID enum
type EventID uint32

// EventID definition
const (
	EventIDTimer         EventID = 10000
	EventIDRequestTrain  EventID = 10011
	EventIDResponseTrain EventID = 10012
	EventIDTrain         EventID = 20001
	EventIDGameFinish    EventID = 20002
)

// IEvent - Event interface
type IEvent interface {
}

// EventRequestTrain - Event definition
type EventRequestTrain struct {
	PlayerID uint64
	BallerID uint64
}

// EventResponseTrain - Event definition
type EventResponseTrain struct {
	ResultCode uint32
	OldLevel   uint32
	NewLevel   uint32
}

// EventTrain - Event Definition
type EventTrain struct {
	PlayerID uint64
	BallerID uint64
	OldLevel uint32
	NewLevel uint32
}
