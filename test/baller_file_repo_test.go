package test

import (
	"clean-arch-example/entity"
	"clean-arch-example/repo"
	"testing"
)

func TestReadWrite(t *testing.T) {
	var baller = entity.Baller{
		ID: 1001,
	}

	var repo = repo.BallerFileRepo{}
	repo.Store(baller)

	var sameBaller = repo.Fetch(1001)
	if sameBaller.ID != baller.ID {
		t.Error("Not same ID")
	}
}
