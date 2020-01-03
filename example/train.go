package example

import (
	"clean-arch-example/usecase"
)

// TrainUcase - Train Usecase
type TrainUcase struct {
	eventManager usecase.IEventManager
	playerRepo   IPlayerRepo
}

// NewTrainUcase - Create new instance
func NewTrainUcase(eventManager usecase.IEventManager, ballerRepo IPlayerRepo) *TrainUcase {
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
	var player = mod.playerRepo.FetchPlayer(playerID)
	var baller = mod.playerRepo.FetchBaller(playerID, ballerID)

	// Change properties
	baller.Level = baller.Level + 1
	player.Exp = player.Exp + 10

	// Store
	mod.playerRepo.StorePlayer(player)
	mod.playerRepo.StoreBaller(baller)

	mod.eventManager.Dispatch(EventIDTrain, EventTrain{
		PlayerID: playerID,
		BallerID: ballerID,
		OldLevel: 1,
		NewLevel: 2,
	})

	var syncData = SyncData{
		Ballers: []*Baller{baller},
		Players: []*Player{player},
	}
	mod.eventManager.Dispatch(EventIDResponseTrain, EventResponseTrain{
		ResultCode: Success,
		Sync:       syncData,
	})
}
