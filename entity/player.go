// Player data
package entity

import "time"

type Player struct {
	ID        uint64
	Name      string
	Exp       uint32
	Level     uint32
	Rank      uint32
	CreatedAt time.Time
}
