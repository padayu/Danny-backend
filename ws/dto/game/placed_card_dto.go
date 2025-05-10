package game

type PlacedCardDto struct {
	Id        int  `json:"id"`
	PositionX int  `json:"position_x"`
	PositionY int  `json:"position_y"`
	Rotation  int  `json:"rotation"`
	Flipped   bool `json:"flipped"`
}
