package example

// ChangedData - Data aggregation that changed in gameplay
type ChangedData struct {
	Players []*Player
	Ballers []*Baller
}

// NewChangedData - Create a new
func NewChangedData() *ChangedData {
	return &ChangedData{}
}

// AddPlayer - add player in sync data
func (data *ChangedData) AddPlayer(player *Player) {
	data.Players = append(data.Players, player)
}

// AddBaller - add baller in sync data
func (data *ChangedData) AddBaller(baller *Baller) {
	data.Ballers = append(data.Ballers, baller)
}
