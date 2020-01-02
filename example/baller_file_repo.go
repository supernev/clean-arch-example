package example

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// FilePath - path to store file
const FilePath = "/tmp"

// BallerFileRepo - A simple implementation of baller repo, with file.
type BallerFileRepo struct {
}

// NewBallerFileRepo - Create a baller repo with json file implementation
func NewBallerFileRepo() IBallerRepo {
	return &BallerFileRepo{}
}

func getFileName(ballerID uint64) string {
	var filename = fmt.Sprintf("%v/Baller_%d.json", FilePath, ballerID)
	return filename
}

// Fetch - fetch a baller from file
// Return baller or nil
func (repo *BallerFileRepo) Fetch(ballerID uint64) Baller {
	var filename = getFileName(ballerID)
	var bytes, _ = ioutil.ReadFile(filename)

	// JSON
	var baller = Baller{}
	json.Unmarshal(bytes, &baller)

	return baller
}

// Store - sotre a baller from file
func (repo *BallerFileRepo) Store(baller Baller) {
	// JSON
	var bytes, _ = json.Marshal(baller)

	var filename = getFileName(baller.ID)
	ioutil.WriteFile(filename, bytes, 0644)
}
