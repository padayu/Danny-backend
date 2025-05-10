package handler

import (
	"danny-backend/logic/game"
	"danny-backend/logic/lobby"
	"danny-backend/ws/api"
	"danny-backend/ws/enums"
	"danny-backend/ws/sender"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

func HandleLobbyCreate(ws *websocket.Conn, server *api.WebSocketServer, message enums.GenericServerBoundMessage) {
	var payload enums.LobbyCreateMessage
	var err = json.Unmarshal(message.Payload, &payload)

	if err != nil {
		log.Println("Unmarshal error:", err)
		return
	}

	code := lobby.ManagerInstance.CreateLobby(payload.Host)

	if code == "" {
		sender.SendStatusMessage(ws, server, 500, "Failed to create lobby")
		return
	}

	sender.SendLobbyCreatedMessage(ws, server, code)
}

func HandleLobbyJoin(ws *websocket.Conn, server *api.WebSocketServer, message enums.GenericServerBoundMessage) {
	var payload enums.LobbyJoinMessage
	var err = json.Unmarshal(message.Payload, &payload)

	if err != nil {
		log.Println("Unmarshal error:", err)
		return
	}

	lobbyInstance := lobby.ManagerInstance.GetLobby(payload.Code)
	if lobbyInstance == nil {
		sender.SendStatusMessage(ws, server, 404, "Lobby not found")
		return
	}

	joinStatus := lobbyInstance.AddMember(payload.Name, ws)
	if !joinStatus {
		sender.SendStatusMessage(ws, server, 500, "Name already taken")
		return
	}

	lobbyInstance.BroadcastState()
}

func HandleLobbyMemberUpdate(ws *websocket.Conn, server *api.WebSocketServer, message enums.GenericServerBoundMessage) {
	var payload enums.LobbyMemberUpdateMessage
	var err = json.Unmarshal(message.Payload, &payload)

	if err != nil {
		log.Println("Unmarshal error:", err)
		return
	}

	lobbyInstance := lobby.ManagerInstance.GetLobby(payload.Code)

	if lobbyInstance == nil {
		sender.SendStatusMessage(ws, server, 404, "Lobby not found")
		return
	}

	errBool := lobbyInstance.UpdateMember(lobby.MemberFromDto(payload.Member, ws))
	if !errBool {
		sender.SendStatusMessage(ws, server, 500, "Failed to update member")
		return
	}

	lobbyInstance.BroadcastState()
}

func HandleLobbyKick(ws *websocket.Conn, server *api.WebSocketServer, message enums.GenericServerBoundMessage) {
	var payload enums.LobbyKickMessage
	var err = json.Unmarshal(message.Payload, &payload)

	if err != nil {
		log.Println("Unmarshal error:", err)
		return
	}

	lobbyInstance := lobby.ManagerInstance.GetLobby(payload.Code)
	if lobbyInstance == nil {
		sender.SendStatusMessage(ws, server, 404, "Lobby not found")
		return
	}

	deletedMembers := lobbyInstance.RemoveMember(payload.Name)
	if len(deletedMembers) == 0 {
		sender.SendStatusMessage(ws, server, 404, "Member not found")
		return
	}

	for _, m := range deletedMembers {
		sender.SendLobbyKickedMessage(m.Connection, server)
	}

	if len(lobbyInstance.Members) == 0 {
		lobby.ManagerInstance.DeleteLobby(payload.Code)
		return
	}

	lobbyInstance.BroadcastState()
}

func HandleLobbyLeave(ws *websocket.Conn, server *api.WebSocketServer, message enums.GenericServerBoundMessage) {
	var payload enums.LobbyKickMessage
	var err = json.Unmarshal(message.Payload, &payload)

	if err != nil {
		log.Println("Unmarshal error:", err)
		return
	}

	lobbyInstance := lobby.ManagerInstance.GetLobby(payload.Code)
	if lobbyInstance == nil {
		sender.SendStatusMessage(ws, server, 404, "Lobby not found")
		return
	}

	deletedMembers := lobbyInstance.RemoveMember(payload.Name)
	if len(deletedMembers) == 0 {
		sender.SendStatusMessage(ws, server, 404, "Member not found")
		return
	}

	sender.SendStatusMessage(ws, server, 200, "Left lobby")

	if len(lobbyInstance.Members) == 0 {
		lobby.ManagerInstance.DeleteLobby(payload.Code)
		return
	}

	lobbyInstance.BroadcastState()
}

func HandleLobbyStartGame(ws *websocket.Conn, server *api.WebSocketServer, message enums.GenericServerBoundMessage) {
	var payload enums.LobbyStartGameMessage
	var err = json.Unmarshal(message.Payload, &payload)

	if err != nil {
		log.Println("Unmarshal error:", err)
		return
	}

	lobbyInstance := lobby.ManagerInstance.GetLobby(payload.Code)
	if lobbyInstance == nil {
		sender.SendStatusMessage(ws, server, 404, "Lobby not found")
		return
	}

	lobbyInstance.BroadcastStart()
	game.ManagerInstance.CreateGame(*lobbyInstance)
	lobby.ManagerInstance.DeleteLobby(lobbyInstance.Code)
}
