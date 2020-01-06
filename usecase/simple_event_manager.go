package usecase

// SimpleEventManager - A simple manager deal with events
type SimpleEventManager struct {
	handlers map[EventID][]func(IEventContext, IEvent)
}

// NewSimpleEventManager - Function to get new instance
func NewSimpleEventManager() *SimpleEventManager {
	return &SimpleEventManager{
		handlers: map[EventID][]func(IEventContext, IEvent){},
	}
}

// Register - Register function to handle a event
func (m *SimpleEventManager) Register(evid EventID, fn func(ctx IEventContext, ev IEvent)) {
	m.handlers[evid] = append(m.handlers[evid], fn)
}

// Dispatch - a event
func (m *SimpleEventManager) Dispatch(ctx IEventContext, evid EventID, ev IEvent) {
	var li, ok = m.handlers[evid]

	if ok {
		for _, fn := range li {
			fn(ctx, ev)
		}
	}
}
