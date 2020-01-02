package example

// IPlayerRepo - Baller repo
type IPlayerRepo interface {
	FetchPlayer(playerID uint64) *Player
	StorePlayer(player *Player)

	FetchBaller(playerID uint64, ballerID uint64) *Baller
	StoreBaller(baller *Baller)
}
