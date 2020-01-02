package example

import (
	"clean-arch-example/usecase"
)

// TrainUcase - Train Usecase
type TrainUcase struct {
	EventManager usecase.IEventManager
	BallerRepo   IBallerRepo
}

// Init method
func (mod *TrainUcase) Init() {
	mod.EventManager.Register(EventIDRequestTrain, mod.handleRequestTrain)
}

// Run method
func (mod *TrainUcase) Run() {
}

// Destroy method
func (mod *TrainUcase) Destroy() {

}

func (mod *TrainUcase) handleRequestTrain(ev usecase.IEvent) {
	var args = ev.(EventRequestTrain)
	var ballerID = args.BallerID
	var playerID = args.PlayerID

	// Fetch
	var baller = mod.BallerRepo.Fetch(ballerID)

	var oldLevel = baller.Level
	// Change properties
	baller.Level = baller.Level + 1
	var newLevel = baller.Level

	// Store
	mod.BallerRepo.Store(baller)

	mod.EventManager.Dispatch(EventIDBallerChanged, EventBallerChanged{
		Ballers: []Baller{baller},
	})

	mod.EventManager.Dispatch(EventIDTrain, EventTrain{
		PlayerID: playerID,
		BallerID: ballerID,
		OldLevel: 1,
		NewLevel: 2,
	})

	mod.EventManager.Dispatch(EventIDResponseTrain, EventResponseTrain{
		ResultCode: Success,
		OldLevel:   oldLevel,
		NewLevel:   newLevel,
	})
}
