package example

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// filePath - path to store file
const filePath = "/tmp"

const prefixPlayer = "Player"
const prefixBaller = "Baller"

// PlayerFileRepo - A simple implementation of player owned data, with file.
type PlayerFileRepo struct {
}

// NewPlayerFileRepo - Create a baller repo with json file implementation
func NewPlayerFileRepo() *PlayerFileRepo {
	return &PlayerFileRepo{}
}

func getPlayerFileName(playerID uint64) string {
	var filename = fmt.Sprintf("%v/%v_%d.json", filePath, prefixPlayer, playerID)
	return filename
}

func getBallerFileName(playerID uint64, ballerID uint64) string {
	var filename = fmt.Sprintf("%v/%v_%d_%d.json", filePath, prefixPlayer, playerID, ballerID)
	return filename
}

// FetchPlayer - fetch a baller from file
func (repo *PlayerFileRepo) FetchPlayer(playerID uint64) *Player {
	var filename = getPlayerFileName(playerID)
	var bytes, _ = ioutil.ReadFile(filename)

	// JSON
	var player = Player{}
	json.Unmarshal(bytes, &player)

	return &player
}

// StorePlayer - sotre a baller from file
func (repo *PlayerFileRepo) StorePlayer(player *Player) {
	// JSON
	var bytes, _ = json.Marshal(player)

	var filename = getPlayerFileName(player.ID)
	ioutil.WriteFile(filename, bytes, 0644)
}

// FetchBaller - fetch a baller from file
func (repo *PlayerFileRepo) FetchBaller(playerID uint64, ballerID uint64) *Baller {
	var filename = getBallerFileName(playerID, ballerID)
	var bytes, _ = ioutil.ReadFile(filename)

	// JSON
	var baller = Baller{}
	json.Unmarshal(bytes, &baller)

	return &baller
}

// StoreBaller - sotre a baller from file
func (repo *PlayerFileRepo) StoreBaller(baller *Baller) {
	// JSON
	var bytes, _ = json.Marshal(baller)

	var filename = getBallerFileName(baller.PlayerID, baller.ID)
	ioutil.WriteFile(filename, bytes, 0644)
}
