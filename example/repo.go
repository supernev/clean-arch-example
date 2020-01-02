package example

// IBallerRepo - Baller repo
type IBallerRepo interface {
	Fetch(ballerID uint64) Baller
	Store(baller Baller)
}
