package game

type PlayerCommonInfoDto struct {
	Name         string `json:"name"`
	HandDeckSize int    `json:"hand_deck_size"`
	IsActive     bool   `json:"is_active"`
	IsDeciding   bool   `json:"is_deciding"`
	VotedWord    string `json:"voted_word"`
}
