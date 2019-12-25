package test

import (
	"clean-arch-example/usecase"
	"testing"
)

func TestRequestTrain(t *testing.T) {
	var mod = usecase.Train{}
	mod.Init()

	var manager = usecase.GetEventManager()
	var ev = usecase.EventRequestTrain{
		PlayerID: 1,
		BallerID: 2,
	}
	manager.Dispatch(usecase.EventIDRequestTrain, ev)

}
