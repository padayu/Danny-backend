package game

import "danny-backend/ws/dto/game"

type Stage int

const (
	CardPlacement Stage = iota
	VotingWord
	VotingPlayer
	EndGame
)

func (s Stage) ToDto() game.Stage {
	switch s {
	case CardPlacement:
		return game.CardPlacement
	case VotingWord:
		return game.VotingWord
	case VotingPlayer:
		return game.VotingPlayer
	case EndGame:
		return game.EndGame
	}

	return "unknown"
}
