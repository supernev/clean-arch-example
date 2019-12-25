package entity

import "time"

// BallerRepo is interface for get/save Baller entity
type BallerRepo interface {
	Get(id uint64) *Baller
	GetList(playerID uint64) []*Baller
	Store(baller *Baller)
	StoreMany(ballers []*Baller)
}

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
