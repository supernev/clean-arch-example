package usecase

import "clean-arch-example/entity"

// IBallerRepo - Baller repo
type IBallerRepo interface {
	Fetch(ballerID uint64) entity.Baller
	Store(baller entity.Baller)
}
