package game

type Stage string

const (
	CardPlacement Stage = "card_placement"
	VotingWord    Stage = "voting_word"
	VotingPlayer  Stage = "voting_player"
	EndGame       Stage = "end_game"
)
