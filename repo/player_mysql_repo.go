package repo

import (
	"clean-arch-example/entity"
)

type BallerMySQLRepo struct {
}

func (repo BallerMySQLRepo) Get(ballerID uint64) *entity.Baller {
	return nil
}
