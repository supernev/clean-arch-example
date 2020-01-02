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
	var ballerRepo = example.NewBallerFileRepo()

	var modTrain = example.NewTrainUcase(eventManager, ballerRepo)

	var app = usecase.NewApplicationSimple()
	app.AddModule(modTrain)

	app.Start()
}
