package entity

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

// CreateBaller public function
func CreateBaller(playerID uint64) *Baller {
	var baller = Baller{
		ID:        1,
		PlayerID:  playerID,
		ConfigID:  10001,
		Exp:       0,
		Level:     1,
		CreatedAt: time.Now(),
	}

	return &baller
}
