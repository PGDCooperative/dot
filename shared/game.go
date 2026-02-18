package shared

type Game struct {
	players  []*Player
	Factions []*Faction `json:"factions"`
	Tiles    []*Tile    `json:"tiles"`
	Entities []*Entity  `json:"entities"`
}

func NewGame() Game {
	return Game{}
}
