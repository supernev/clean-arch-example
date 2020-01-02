package example

import "time"

// Baller entity
type Baller struct {
	ID        uint64
	PlayerID  uint64
	ConfigID  uint32
	Exp       uint32
	Level     uint32
	CreatedAt time.Time
}
