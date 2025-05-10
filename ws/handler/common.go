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

func HandleRestoreConnection(ws *websocket.Conn, server *api.WebSocketServer, message enums.GenericServerBoundMessage) {
	var payload enums.RestoreConnectionMessage
	var err = json.Unmarshal(message.Payload, &payload)

	if err != nil {
		log.Println("Unmarshal error:", err)
		return
	}

	gameInstance := game.ManagerInstance.GetGame(payload.Code)
	if gameInstance != nil {
		result := gameInstance.RestoreConnection(payload.Name, ws)

		if !result {
			sender.SendStatusMessage(ws, server, 404, "No such player in game")
			return
		}

		sender.SendStatusMessage(ws, server, 200, "Connection restored")
		return
	}

	lobbyInstance := lobby.ManagerInstance.GetLobby(payload.Code)
	if lobbyInstance != nil {
		result := lobbyInstance.RestoreConnection(payload.Name, ws)

		if !result {
			sender.SendStatusMessage(ws, server, 404, "No such member in lobby")
			return
		}

		sender.SendStatusMessage(ws, server, 200, "Connection restored")
		return
	}

	sender.SendStatusMessage(ws, server, 404, "No such game or lobby")
}
