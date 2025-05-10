package enums

import (
	"danny-backend/ws/dto/game"
	"danny-backend/ws/dto/lobby"
	"encoding/json"
)

type ClientBoundMessage string

const (
	StatusCode       ClientBoundMessage = "status_code"
	LobbyCreated     ClientBoundMessage = "lobby_created"
	LobbyInfo        ClientBoundMessage = "lobby_info"
	LobbyKicked      ClientBoundMessage = "lobby_kicked"
	LobbyGameStarted ClientBoundMessage = "lobby_game_started"

	GameInfo ClientBoundMessage = "game_info"
	GameChat ClientBoundMessage = "game_chat"

	GameFinalScores ClientBoundMessage = "game_final_scores"
)

type GenericClientBoundMessage struct {
	Type    ClientBoundMessage `json:"type"`
	Payload json.RawMessage    `json:"payload"`
}

type StatusCodeMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type LobbyCreatedMessage struct {
	Code string `json:"code"`
}

type LobbyInfoMessage struct {
	State lobby.StateDto `json:"state"`
}

type LobbyKickedMessage struct{}

type LobbyStartedMessage struct{}

type GameInfoMessage struct {
	State game.StateDto `json:"state"`
}

type GameChatMessage struct {
	Message game.ChatMessageDto `json:"message"`
}

type GameFinalScoresMessage struct {
	Scores []game.ScoreboardEntryDto `json:"scores"`
}
