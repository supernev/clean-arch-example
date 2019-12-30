package repo

import (
	"clean-arch-example/entity"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Path - path to store file
const PATH = "/tmp"

// BallerFileRepo - A simple implementation of baller repo, with file.
type BallerFileRepo struct {
}

func getFileName(ballerID uint64) string {
	var filename = fmt.Sprintf("%v/Baller_%d.json", PATH, ballerID)
	return filename
}

// Fetch - fetch a baller from file
// Return baller or nil
func (repo *BallerFileRepo) Fetch(ballerID uint64) entity.Baller {
	var filename = getFileName(ballerID)
	var bytes, _ = ioutil.ReadFile(filename)

	// JSON
	var baller = entity.Baller{}
	json.Unmarshal(bytes, &baller)

	return baller
}

// Store - sotre a baller from file
func (repo *BallerFileRepo) Store(baller entity.Baller) {
	// JSON
	var bytes, _ = json.Marshal(baller)

	var filename = getFileName(baller.ID)
	ioutil.WriteFile(filename, bytes, 0644)
}
