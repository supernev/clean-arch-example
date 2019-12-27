package usecase

import (
	"fmt"
)

// TrainUcase - Train Usecase
type TrainUcase struct {
	eventManager *EventManager
}

// Init method
func (mod *TrainUcase) Init(eventManager *EventManager) {
	mod.eventManager = eventManager
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
