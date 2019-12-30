package usecase

import "container/list"

// EventID enum
type EventID uint32

// EventID definition
const (
	// Timer
	EventIDTimer EventID = 10000

	// Client events
	EventIDRequestTrain  EventID = 10011
	EventIDResponseTrain EventID = 10012

	// Gameplay events
	EventIDTrain      EventID = 20001
	EventIDGameFinish EventID = 20002

	// Entity change events
	EventIDBallerChanged EventID = 90001
)

// IEvent - Event interface
type IEvent interface {
}

// IEventManager - Event manager interface
type IEventManager interface {
	Register(evid EventID, fn func(ev IEvent))
	Dispatch(evid EventID, ev IEvent)
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

// EventBallerChanged - Event when data change
type EventBallerChanged struct {
	Ballers list.List
}
