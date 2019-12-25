package usecase

import (
	"fmt"
)

// Train usecase
type Train struct {
}

// Init method
func (mod Train) Init() {
	var man = GetEventManager()
	man.Register(EventIDRequestTrain, mod.handleRequestTrain)
}

// Run method
func (mod Train) Run() {

}

// Destroy method
func (mod Train) Destroy() {

}

func (mod Train) handleRequestTrain(ev IEvent) {
	var req = ev.(EventRequestTrain)
	var ballerID = req.BallerID
	var playerID = req.BallerID
	fmt.Println("PlayerID=", playerID, "BallerID=", ballerID)

	// var tx = repo.StartTransaction()
	// var baller = tx.GetBaller(ballerID)
	// baller.lv += 1
	// tx.commit()

	var man = GetEventManager()

	man.Dispatch(EventIDTrain, EventTrain{
		PlayerID: playerID,
		BallerID: ballerID,
		OldLevel: 1,
		NewLevel: 2,
	})

	man.Dispatch(EventIDResponseTrain, EventResponseTrain{
		ResultCode: 0,
		OldLevel:   1,
		NewLevel:   2,
	})
}
