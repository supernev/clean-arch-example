package usecase

var manager = EventManager{
	handlers: map[EventID]func(IEvent){},
}

// EventManager - A manager deal with events
type EventManager struct {
	handlers map[EventID]func(IEvent)
}

// GetEventManager - Get singeleton
func GetEventManager() *EventManager {
	return &manager
}

// Register - Register function to handle a event
func (m EventManager) Register(evid EventID, fn func(ev IEvent)) {
	m.handlers[evid] = fn
}

// Dispatch - a event
func (m EventManager) Dispatch(evid EventID, ev IEvent) {
	var fn = m.handlers[evid]
	if fn != nil {
		fn(ev)
	}
}
