package test

import (
	"clean-arch-example/example"
	"testing"
)

func TestReadWrite(t *testing.T) {
	var baller = example.Baller{
		ID:       1001,
		PlayerID: 10001,
	}

	var repo = example.NewPlayerFileRepo()
	repo.StoreBaller(&baller)

	var sameBaller = repo.FetchBaller(10001, 1001)
	if sameBaller.ID != baller.ID {
		t.Error("Not same ID")
	}
}
