package game

import (
	"danny-backend/ws/dto/game"
	"danny-backend/ws/sender"
	"github.com/gorilla/websocket"
)

func (g *Game) SendServerMessage(playerInstance *Player, message string) {
	sender.SendChatMessage(playerInstance.Connection, ManagerInstance.Server, game.ChatMessageDto{
		Sender: "Server",
		Text:   message,
	})
}

func (g *Game) BroadcastServerMessage(message string) {
	for _, playerInstance := range g.players {
		g.SendServerMessage(playerInstance, message)
	}
}

func (g *Game) SendState(playerInstance *Player) {
	sender.SendGameStateMessage(playerInstance.Connection, ManagerInstance.Server, g.ToDto(playerInstance))
}

func (g *Game) BroadcastState() {
	for _, playerInstance := range g.players {
		g.SendState(playerInstance)
	}
}

func (g *Game) RestoreConnection(name string, connection *websocket.Conn) bool {
	for _, p := range g.players {
		if p.Name == name {
			p.Connection = connection
			g.SendState(p)
			return true
		}
	}

	return false
}
