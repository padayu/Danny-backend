package game

import (
	"danny-backend/ws/dto/game"
	"danny-backend/ws/sender"
)

func (g *Game) endGame() {
	g.BroadcastServerMessage("Игра окончена")
	g.stage = EndGame

	g.BroadcastState()

	scores := make([]game.ScoreboardEntryDto, 0)

	for _, p := range g.players {
		won := calculateOutcome(g, p)

		scores = append(scores, game.ScoreboardEntryDto{
			Name: p.Name,
			Won:  won,
		})
	}

	for _, p := range g.players {
		sender.SendGameFinalScoresMessage(p.Connection, ManagerInstance.Server, scores)
	}

	g.BroadcastServerMessage("Игра окончена! Победил " + scores[0].Name + "!")
}

func calculateOutcome(g *Game, p *Player) bool {
	return (p.IsDanny && g.scoreD >= 3) || (!p.IsDanny && g.scoreA >= 6)
}
