package sender

import (
	"danny-backend/ws/api"
	"danny-backend/ws/dto/lobby"
	"danny-backend/ws/enums"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

func SendStatusMessage(ws *websocket.Conn, server *api.WebSocketServer, status int, message string) {
	var payload, err = json.Marshal(enums.StatusCodeMessage{Status: status, Message: message})

	if err != nil {
		log.Println("Marshal error:", err)
		return
	}

	var data = enums.GenericClientBoundMessage{
		Type:    enums.StatusCode,
		Payload: json.RawMessage(payload),
	}

	server.MessageSender(ws, data)
}

func SendLobbyCreatedMessage(ws *websocket.Conn, server *api.WebSocketServer, code string) {
	var payload, err = json.Marshal(enums.LobbyCreatedMessage{Code: code})

	if err != nil {
		log.Println("Marshal error:", err)
		return
	}

	var data = enums.GenericClientBoundMessage{
		Type:    enums.LobbyCreated,
		Payload: json.RawMessage(payload),
	}

	server.MessageSender(ws, data)
}

func SendLobbyInfoMessage(ws *websocket.Conn, server *api.WebSocketServer, state lobby.StateDto) {
	var payload, err = json.Marshal(enums.LobbyInfoMessage{State: state})

	if err != nil {
		log.Println("Marshal error:", err)
		return
	}

	var data = enums.GenericClientBoundMessage{
		Type:    enums.LobbyInfo,
		Payload: json.RawMessage(payload),
	}

	server.MessageSender(ws, data)
}

func SendLobbyKickedMessage(ws *websocket.Conn, server *api.WebSocketServer) {
	var data = enums.GenericClientBoundMessage{
		Type:    enums.LobbyKicked,
		Payload: nil,
	}

	server.MessageSender(ws, data)
}

func SendLobbyStartedMessage(ws *websocket.Conn, server *api.WebSocketServer) {
	var data = enums.GenericClientBoundMessage{
		Type:    enums.LobbyGameStarted,
		Payload: nil,
	}

	server.MessageSender(ws, data)
}
