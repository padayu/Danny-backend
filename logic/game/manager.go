package game

import (
	"danny-backend/logic/lobby"
	"danny-backend/ws/api"
	"math/rand"
)

type Manager struct {
	games  map[string]*Game
	Server *api.WebSocketServer
}

var ManagerInstance Manager

func InitManager(server *api.WebSocketServer) {
	ManagerInstance = Manager{
		games:  make(map[string]*Game),
		Server: server,
	}
}

func (m *Manager) CreateGame(lobby lobby.Lobby) {
	var players = make([]*Player, 0)
	DannyIndex := rand.Intn(len(lobby.Members))
	for index, m := range lobby.Members {
		players = append(players, &Player{
			Name:       m.Name,
			Connection: m.Connection,
			IsDanny:    index == DannyIndex,
		})
	}

	m.games[lobby.Code] = &Game{
		code:    lobby.Code,
		players: players,
	}

	m.games[lobby.Code].StartGame()
}

func (m *Manager) GetGame(code string) *Game {
	game, ok := m.games[code]
	if !ok {
		return nil
	}

	return game
}
