package game

type StateDto struct {
	Players       []PlayerCommonInfoDto `json:"players"`
	Player        PlayerInfoDto         `json:"player"`
	Stage         Stage                 `json:"stage"`
	CardPlacement CardPlacementDto      `json:"card_placement"`
	WordOptions   []string              `json:"word_options"`
	Word          string                `json:"word"`
	ScoreA        int                   `json:"score_a"`
	ScoreD        int                   `json:"score_d"`
}
