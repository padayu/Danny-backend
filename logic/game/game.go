package game

import (
	"danny-backend/logic/game/card"
	"danny-backend/ws/dto/game"
	"math/rand"
)

type Game struct {
	code                     string
	players                  []*Player
	stage                    Stage
	mainDeck                 []card.Memory
	cardPlacement            []PlacedCard
	activePersonalityIndex   int
	decidingPersonalityIndex int
	playerRotationIndexStart int
	chatHistory              []ChatMessage
	scoreD                   int
	scoreA                   int
	wordOptions              []string
	word                     string
	endGameConditionMet      bool
}

func (g *Game) StartGame() {
	g.activePersonalityIndex = 0
	g.decidingPersonalityIndex = len(g.players) - 1
	g.BroadcastServerMessage("Игра началась!")

	rand.Shuffle(len(g.players), func(i, j int) {
		g.players[i], g.players[j] = g.players[j], g.players[i]
	})

	g.mainDeck = card.GenerateMainDeck()

	g.StartTurn()
}

func (g *Game) StartTurn() {
	g.stage = CardPlacement
	g.wordOptions = GenerateWordOptions()
	wordIndex := rand.Intn(len(g.wordOptions))
	g.word = g.wordOptions[wordIndex]
	g.BroadcastServerMessage("Объяснение слова началось")

	activePlayer := g.GetActivePlayer()

	if activePlayer != nil {
		g.BroadcastServerMessage("Ход игрока " + activePlayer.Name)
		g.DrawCards(activePlayer)
	}

	g.BroadcastState()
}

func (g *Game) StartVoteWord() {
	g.stage = VotingWord
	g.BroadcastServerMessage("Угадывание слова началось")

	decidingPlayer := g.GetDecidingPlayer()

	if decidingPlayer != nil {
		g.BroadcastServerMessage("Последнее слово за игроком " + decidingPlayer.Name)
	}

	g.BroadcastState()
}

func (g *Game) ReturnCardToMainDeck(card card.Memory) {
	g.mainDeck = append(g.mainDeck, card)
}

func (g *Game) EndTurn() {
	if g.endGameConditionMet {
		g.endGame()
		return
	}

	g.activePersonalityIndex++
	if g.activePersonalityIndex >= len(g.players) {
		g.activePersonalityIndex = 0
	}

	g.decidingPersonalityIndex++
	if g.decidingPersonalityIndex >= len(g.players) {
		g.decidingPersonalityIndex = 0
	}

	g.StartTurn()
}

func (g *Game) GetActivePlayer() *Player {
	return g.players[g.activePersonalityIndex]
}

func (g *Game) GetDecidingPlayer() *Player {
	return g.players[g.decidingPersonalityIndex]
}

func (g *Game) GetPlayer(name string) *Player {
	for _, p := range g.players {
		if p.Name == name {
			return p
		}
	}

	return nil
}

func (g *Game) VotePlayer(voter string, suspect string) bool {

	voterInstance := g.GetPlayer(voter)
	suspectInstance := g.GetPlayer(suspect)
	if voterInstance == nil || suspectInstance == nil {
		return false
	}

	voterInstance.VotePlayer(suspect)
	return true
}

func (g *Game) VoteWord(voter string, word string) bool {

	voterInstance := g.GetPlayer(voter)
	if voterInstance == nil {
		return false
	}

	voterInstance.VoteWord(word)
	if g.GetDecidingPlayer() == voterInstance {
		g.DeciderChooseWord(word)
	}
	return true
}

func (g *Game) DeciderChooseWord(word string) {
	if word == g.word {
		g.scoreA++
	} else {
		g.scoreD++
	}
	if g.scoreA >= 6 || g.scoreD >= 3 {
		g.endGameConditionMet = true
	}
	g.EndTurn()
}

func (g *Game) AddChatMessage(name string, text string) {
	g.chatHistory = append(g.chatHistory, ChatMessage{name, text})
}

func (g *Game) UpdateCardPlacement(placement []PlacedCard) {
	g.cardPlacement = placement
}

func (g *Game) DrawCards(player *Player) {
	player.ReceiveHand(g.mainDeck[:7])
	g.ShuffleDeck()
}

func (g *Game) ShuffleDeck() {
	rand.Shuffle(len(g.mainDeck), func(i, j int) {
		g.mainDeck[i], g.mainDeck[j] = g.mainDeck[j], g.mainDeck[i]
	})
}

func (g *Game) ToDto(playerInstance *Player) game.StateDto {
	return g.generateStateDto(playerInstance)
}
