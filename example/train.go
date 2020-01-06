package example

import (
	"clean-arch-example/usecase"
)

// TrainUcase - Train Usecase
type TrainUcase struct {
	eventManager usecase.IEventManager
}

// NewTrainUcase - Create new instance
func NewTrainUcase(eventManager usecase.IEventManager) *TrainUcase {
	return &TrainUcase{eventManager}
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

func (mod *TrainUcase) handleRequestTrain(ctx usecase.IEventContext, ev usecase.IEvent) {
	var args = ev.(EventRequestTrain)
	var ballerID = args.BallerID
	var playerID = args.PlayerID

	var playerRepo = ctx.Get(ContextKeyPlayerRepo).(IPlayerRepo)

	// Fetch
	var player = playerRepo.FetchPlayer(playerID)
	var baller = playerRepo.FetchBaller(playerID, ballerID)

	// Change properties
	baller.Level = baller.Level + 1
	player.Exp = player.Exp + 10

	// Store
	playerRepo.StorePlayer(player)
	playerRepo.StoreBaller(baller)

	var changedData = NewChangedData()
	changedData.AddPlayer(player)
	changedData.AddBaller(baller)

	ctx.Set(ContextKeyChangedData, changedData)

	// Dispatch a "train" event for other modules to handle
	mod.eventManager.Dispatch(ctx, EventIDTrain, EventTrain{
		PlayerID: playerID,
		BallerID: ballerID,
		OldLevel: 1,
		NewLevel: 2,
	})

	// Dispatch a "response" event for other modules to response to client
	mod.eventManager.Dispatch(ctx, EventIDResponseTrain, EventResponseTrain{
		ResultCode:  Success,
		ChangedData: changedData,
	})
}
