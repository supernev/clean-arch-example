package main

import (
	"clean-arch-example/example"
	"clean-arch-example/usecase"
	"fmt"
)

func main() {
	fmt.Println("Starting example")
	defer fmt.Println("Stopping example")

	var eventManager usecase.IEventManager = usecase.GetSimpleEventManager()
	var ballerRepo example.IBallerRepo = example.NewBallerFileRepo()

	var modTrain = example.TrainUcase{EventManager: eventManager, BallerRepo: ballerRepo}

	var app = usecase.ApplicationSimple{}
	app.AddModule(&modTrain)

	app.Start()
}
