package lobby

import (
	"danny-backend/ws/api"
	"math/rand"
	"time"
)

const (
	codeLength = 4
)

type Manager struct {
	Lobbies map[string]*Lobby
	Server  *api.WebSocketServer
}

var ManagerInstance Manager

func InitManager(server *api.WebSocketServer) {
	ManagerInstance = Manager{
		Lobbies: make(map[string]*Lobby),
		Server:  server,
	}
}

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"

func stringWithCharset(length int, charset string) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (m *Manager) GenerateLobbyCode() string {
	for {
		code := stringWithCharset(codeLength, charset)
		if _, ok := m.Lobbies[code]; !ok {
			return code
		}
	}
}

func (m *Manager) CreateLobby(host string) string {
	code := m.GenerateLobbyCode()
	m.Lobbies[code] = &Lobby{
		Code:    code,
		Host:    host,
		Members: []*Member{},
	}

	return code
}

func (m *Manager) GetLobby(code string) *Lobby {
	lobby, ok := m.Lobbies[code]
	if !ok {
		return nil
	}

	return lobby
}

func (m *Manager) DeleteLobby(code string) {
	delete(m.Lobbies, code)
}
