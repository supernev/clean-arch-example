package usecase

import (
	"clean-arch-example/entity"
)

// TrainUcase - Train Usecase
type TrainUcase struct {
	eventManager IEventManager
	ballerRepo   IBallerRepo
	agent        IAgent
}

// Init method
func (mod *TrainUcase) Init(
	eventManager IEventManager,
	ballerRepo IBallerRepo,
	agent IAgent,
) {
	mod.eventManager = eventManager
	mod.ballerRepo = ballerRepo
	mod.agent = agent

	mod.agent.OnRequestTrain(mod.handleRequestTrain)
}

// Run method
func (mod *TrainUcase) Run() {

}

// Destroy method
func (mod *TrainUcase) Destroy() {

}

func (mod *TrainUcase) handleRequestTrain(ballerID uint64, playerID uint64) {
	// Fetch
	var baller = mod.ballerRepo.Fetch(ballerID)

	var oldLevel = baller.Level
	// Change properties
	baller.Level = baller.Level + 1
	var newLevel = baller.Level

	// Store
	mod.ballerRepo.Store(baller)

	mod.eventManager.Dispatch(EventIDBallerChanged, EventBallerChanged{
		Ballers: []entity.Baller{baller},
	})

	mod.eventManager.Dispatch(EventIDTrain, EventTrain{
		PlayerID: playerID,
		BallerID: ballerID,
		OldLevel: 1,
		NewLevel: 2,
	})

	mod.agent.ResponseTrain(Success, oldLevel, newLevel)
}
