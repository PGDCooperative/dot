package shared

type TType int

const (
	AIR TType = iota
	STONE
)

type Tile struct {
	HP        int `json:"hp"`
	ID        int `json:"id"`
	FactionID int `json:"factionid"`
	TType     `json:"type"`
}
