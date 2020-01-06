package usecase

// EventID enum
type EventID = uint32

// IEvent - Event interface
type IEvent interface {
}

// IEventContext - context the event happens
type IEventContext interface {
	Get(name string) interface{}
	Set(name string, value interface{})
}

// IEventManager - Event manager interface
type IEventManager interface {
	Register(evid EventID, fn func(ctx IEventContext, ev IEvent))
	Dispatch(ctx IEventContext, evid EventID, ev IEvent)
}
