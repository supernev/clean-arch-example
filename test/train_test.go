package test

import (
	"clean-arch-example/entity"
	"clean-arch-example/repo"
	"clean-arch-example/usecase"
	"testing"
	"time"
)

func TestRequestTrain(t *testing.T) {
	var mod = usecase.TrainUcase{}
	var manager = usecase.GetSimpleEventManager()
	var repo = repo.BallerFileRepo{}
	mod.Init(manager, &repo)

	// Setup data
	var baller = entity.Baller{
		ID:        1001,
		PlayerID:  1,
		ConfigID:  654321,
		Exp:       0,
		Level:     1,
		CreatedAt: time.Now(),
	}
	repo.Store(baller)

	var ev = usecase.EventRequestTrain{
		PlayerID: 1,
		BallerID: 1001,
	}
	manager.Dispatch(usecase.EventIDRequestTrain, ev)

	var ballerAgain = repo.Fetch(1001)
	if ballerAgain.ID != 1001 {
		t.Error("Wrong ID")
	}
	if ballerAgain.Level != 2 {
		t.Error("Wrong Level")
	}
}
