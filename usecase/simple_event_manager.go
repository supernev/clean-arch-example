package usecase

// SimpleEventManager - A simple manager deal with events
type SimpleEventManager struct {
	handlers map[EventID]func(IEvent)
}

// GetSimpleEventManager - Static function to get
func GetSimpleEventManager() *SimpleEventManager {
	return &SimpleEventManager{
		handlers: map[EventID]func(IEvent){},
	}
}

// Register - Register function to handle a event
func (m *SimpleEventManager) Register(evid EventID, fn func(ev IEvent)) {
	m.handlers[evid] = fn
}

// Dispatch - a event
func (m *SimpleEventManager) Dispatch(evid EventID, ev IEvent) {
	var fn = m.handlers[evid]
	if fn != nil {
		fn(ev)
	}
}
