package enums

import (
	"danny-backend/ws/dto/game"
	"danny-backend/ws/dto/lobby"
	"encoding/json"
)

type ServerBoundMessage string

const (
	RestoreConnection ServerBoundMessage = "restore_connection"

	LobbyCreate       ServerBoundMessage = "lobby_create"
	LobbyJoin         ServerBoundMessage = "lobby_join"
	LobbyMemberUpdate ServerBoundMessage = "lobby_member_update"
	LobbyKick         ServerBoundMessage = "lobby_kick"
	LobbyLeave        ServerBoundMessage = "lobby_leave"
	LobbyStartGame    ServerBoundMessage = "lobby_start_game"

	GameVoteWord          ServerBoundMessage = "game_vote_word"
	GameVotePlayer        ServerBoundMessage = "game_vote_player"
	GameCardPlacement     ServerBoundMessage = "game_card_placement"
	GamePlayerChatMessage ServerBoundMessage = "game_player_chat_message"
)

type GenericServerBoundMessage struct {
	Type    ServerBoundMessage `json:"type"`
	Payload json.RawMessage    `json:"payload"`
}

type RestoreConnectionMessage struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type LobbyCreateMessage struct {
	Host string `json:"host"`
}

type LobbyJoinMessage struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type LobbyMemberUpdateMessage struct {
	Code   string          `json:"code"`
	Member lobby.MemberDto `json:"member"`
}

type LobbyKickMessage struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type LobbyStartGameMessage struct {
	Code string `json:"code"`
}

type GameVoteWordMessage struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Word string `json:"word"`
}

type GamePlayerChatMessageMessage struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

type GameCardPlacementMessage struct {
	Code      string                `json:"code"`
	Name      string                `json:"name"`
	Placement game.CardPlacementDto `json:"placement"`
}

type GameVotePlayerMessage struct {
	Code       string `json:"code"`
	SenderName string `json:"sender_name"`
	VotedName  string `json:"voted_name"`
}
