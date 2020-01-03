package main

import (
	"clean-arch-example/example"
	"clean-arch-example/usecase"
	"fmt"
)

func main() {
	fmt.Println("Starting example")
	defer fmt.Println("Stopping example")

	var eventManager = usecase.NewSimpleEventManager()
	var playerRepo = example.NewPlayerFileRepo()

	var modTrain = example.NewTrainUcase(eventManager, playerRepo)

	var app = usecase.NewSimpleApplication()
	app.AddModule(modTrain)

	app.Start()
}
