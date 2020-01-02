package usecase

// EventID enum
type EventID = uint32

// IEvent - Event interface
type IEvent interface {
}

// IEventManager - Event manager interface
type IEventManager interface {
	Register(evid EventID, fn func(ev IEvent))
	Dispatch(evid EventID, ev IEvent)
}
