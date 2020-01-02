package test

import (
	"clean-arch-example/example"
	"testing"
)

func TestReadWrite(t *testing.T) {
	var baller = example.Baller{
		ID: 1001,
	}

	var repo = example.BallerFileRepo{}
	repo.Store(baller)

	var sameBaller = repo.Fetch(1001)
	if sameBaller.ID != baller.ID {
		t.Error("Not same ID")
	}
}
