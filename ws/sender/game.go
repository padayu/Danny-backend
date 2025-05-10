package sender

import (
	"danny-backend/ws/api"
	"danny-backend/ws/dto/game"
	"danny-backend/ws/enums"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

func SendChatMessage(ws *websocket.Conn, server *api.WebSocketServer, message game.ChatMessageDto) {
	var payload, err = json.Marshal(enums.GameChatMessage{Message: message})

	if err != nil {
		log.Println("Marshal error:", err)
		return
	}

	var data = enums.GenericClientBoundMessage{
		Type:    enums.GameChat,
		Payload: json.RawMessage(payload),
	}

	server.MessageSender(ws, data)
}

func SendGameStateMessage(ws *websocket.Conn, server *api.WebSocketServer, state game.StateDto) {
	var payload, err = json.Marshal(enums.GameInfoMessage{State: state})

	if err != nil {
		log.Println("Marshal error:", err)
		return
	}

	var data = enums.GenericClientBoundMessage{
		Type:    enums.GameInfo,
		Payload: json.RawMessage(payload),
	}

	server.MessageSender(ws, data)
}

func SendGameFinalScoresMessage(ws *websocket.Conn, server *api.WebSocketServer, scores []game.ScoreboardEntryDto) {
	var payload, err = json.Marshal(enums.GameFinalScoresMessage{Scores: scores})

	if err != nil {
		log.Println("Marshal error:", err)
		return
	}

	var data = enums.GenericClientBoundMessage{
		Type:    enums.GameFinalScores,
		Payload: json.RawMessage(payload),
	}

	server.MessageSender(ws, data)
}
