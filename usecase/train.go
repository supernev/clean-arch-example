package usecase

import (
	"container/list"
	"fmt"
)

// TrainUcase - Train Usecase
type TrainUcase struct {
	eventManager IEventManager
	ballerRepo   IBallerRepo
}

// Init method
func (mod *TrainUcase) Init(eventManager IEventManager, ballerRepo IBallerRepo) {
	mod.eventManager = eventManager
	mod.ballerRepo = ballerRepo
	mod.eventManager.Register(EventIDRequestTrain, mod.handleRequestTrain)
}

// Run method
func (mod *TrainUcase) Run() {

}

// Destroy method
func (mod *TrainUcase) Destroy() {

}

func (mod *TrainUcase) handleRequestTrain(ev IEvent) {
	var req = ev.(EventRequestTrain)
	var ballerID = req.BallerID
	var playerID = req.BallerID
	fmt.Println("PlayerID=", playerID, "BallerID=", ballerID)

	// Fetch
	var baller = mod.ballerRepo.Fetch(ballerID)

	// Change properties
	baller.Level = baller.Level + 1

	// Store
	mod.ballerRepo.Store(baller)

	var changedBallers = list.List{}
	changedBallers.PushBack(baller)
	mod.eventManager.Dispatch(EventIDBallerChanged, EventBallerChanged{
		Ballers: changedBallers,
	})

	mod.eventManager.Dispatch(EventIDTrain, EventTrain{
		PlayerID: playerID,
		BallerID: ballerID,
		OldLevel: 1,
		NewLevel: 2,
	})

	mod.eventManager.Dispatch(EventIDResponseTrain, EventResponseTrain{
		ResultCode: 0,
		OldLevel:   1,
		NewLevel:   2,
	})
}
