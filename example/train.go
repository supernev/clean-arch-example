package example

import (
	"clean-arch-example/usecase"
)

// TrainUcase - Train Usecase
type TrainUcase struct {
	eventManager usecase.IEventManager
	ballerRepo   IBallerRepo
}

// NewTrainUcase - Create new instance
func NewTrainUcase(eventManager usecase.IEventManager, ballerRepo IBallerRepo) *TrainUcase {
	return &TrainUcase{eventManager, ballerRepo}
}

// Init method
func (mod *TrainUcase) Init() {
	mod.eventManager.Register(EventIDRequestTrain, mod.handleRequestTrain)
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
	var baller = mod.ballerRepo.Fetch(ballerID)

	var oldLevel = baller.Level
	// Change properties
	baller.Level = baller.Level + 1
	var newLevel = baller.Level

	// Store
	mod.ballerRepo.Store(baller)

	mod.eventManager.Dispatch(EventIDBallerChanged, EventBallerChanged{
		Ballers: []Baller{baller},
	})

	mod.eventManager.Dispatch(EventIDTrain, EventTrain{
		PlayerID: playerID,
		BallerID: ballerID,
		OldLevel: 1,
		NewLevel: 2,
	})

	mod.eventManager.Dispatch(EventIDResponseTrain, EventResponseTrain{
		ResultCode: Success,
		OldLevel:   oldLevel,
		NewLevel:   newLevel,
	})
}
