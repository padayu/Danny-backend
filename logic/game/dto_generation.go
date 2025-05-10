package game

import (
	"danny-backend/ws/dto/game"
)

func (g *Game) generateStateDto(playerInstance *Player) game.StateDto {

	return game.StateDto{
		Players:     g.generatePlayersDto(),
		Player:      g.generatePlayerDto(playerInstance),
		Stage:       g.stage.ToDto(),
		WordOptions: g.wordOptions,
		Word:        g.word,
		ScoreA:      g.scoreA,
		ScoreD:      g.scoreD,
	}
}

func (g *Game) generatePlayersDto() []game.PlayerCommonInfoDto {
	players := make([]game.PlayerCommonInfoDto, 0)
	for _, p := range g.players {
		players = append(players, p.ToCommonInfoDto(g))
	}

	return players
}

func (g *Game) generatePlayerDto(playerInstance *Player) game.PlayerInfoDto {
	return playerInstance.ToInfoDto(g)
}
