package test

import (
	"clean-arch-example/example"
	"clean-arch-example/usecase"
	"testing"
	"time"
)

func TestRequestTrain(t *testing.T) {
	var eventManager = usecase.NewSimpleEventManager()
	var repo = example.NewPlayerFileRepo()
	var mod = example.NewTrainUcase(eventManager, repo)
	mod.Init()

	// Setup data
	var baller = example.Baller{
		ID:        1001,
		PlayerID:  10001,
		ConfigID:  654321,
		Exp:       0,
		Level:     1,
		CreatedAt: time.Now(),
	}
	repo.StoreBaller(&baller)

	var ev = example.EventRequestTrain{
		PlayerID: 10001,
		BallerID: 1001,
	}
	eventManager.Dispatch(example.EventIDRequestTrain, ev)

	var ballerAgain = repo.FetchBaller(10001, 1001)
	if ballerAgain.ID != 1001 {
		t.Error("Wrong ID")
	}
	if ballerAgain.Level != 2 {
		t.Error("Wrong Level")
	}
}
