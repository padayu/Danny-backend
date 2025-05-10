package game

import (
	"danny-backend/logic/game/card"
	"danny-backend/ws/dto/game"
	"github.com/gorilla/websocket"
)

type Player struct {
	Name        string
	Hand        []card.Memory
	IsDanny     bool
	VotedWord   string
	VotedPlayer string
	Connection  *websocket.Conn
}

func (p *Player) GiveCard(card card.Memory) {
	p.Hand = append(p.Hand, card)
}

func (p *Player) VotePlayer(suspect string) {
	p.VotedPlayer = suspect
}

func (p *Player) VoteWord(word string) {
	p.VotedWord = word
}

func (p *Player) ReceiveHand(hand []card.Memory) {
	p.Hand = hand
}

func (p *Player) ToInfoDto(g *Game) game.PlayerInfoDto {
	hand := make([]game.MemoryCardDto, 0)
	for _, b := range p.Hand {
		hand = append(hand, b.ToDto())
	}

	return game.PlayerInfoDto{
		Hand:       hand,
		IsActive:   g.players[g.activePersonalityIndex] == p,
		IsDeciding: g.players[g.decidingPersonalityIndex] == p,
		IsDanny:    p.IsDanny,
	}
}

func (p *Player) ToCommonInfoDto(g *Game) game.PlayerCommonInfoDto {

	return game.PlayerCommonInfoDto{
		Name:         p.Name,
		HandDeckSize: len(p.Hand),
		IsActive:     g.players[g.activePersonalityIndex] == p,
		IsDeciding:   g.players[g.decidingPersonalityIndex] == p,
		VotedWord:    p.VotedWord,
	}
}
