package entity

import "time"

type BallerRepo interface {
	Get(id uint64) *Baller
	GetList(playerID uint64) []*Baller
	Store(baller *Baller)
	StoreMany(ballers []*Baller)
}

type Baller struct {
	ID        uint64
	PlayerID  uint64
	ConfigID  uint32
	Exp       uint32
	Level     uint32
	CreatedAt time.Time
}

func Create(playerID uint64) *Baller {
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
