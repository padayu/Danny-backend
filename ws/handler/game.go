package handler

import (
	"danny-backend/logic/game"
	"danny-backend/ws/api"
	"danny-backend/ws/enums"
	"danny-backend/ws/sender"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

func HandleGameVotePlayer(ws *websocket.Conn, server *api.WebSocketServer, message enums.GenericServerBoundMessage) {
	var payload enums.GameVotePlayerMessage
	var err = json.Unmarshal(message.Payload, &payload)

	if err != nil {
		log.Println("Unmarshal error:", err)
		return
	}

	gameInstance := game.ManagerInstance.GetGame(payload.Code)
	if gameInstance == nil {
		sender.SendStatusMessage(ws, server, 404, "Game not found")
		return
	}

	ok := gameInstance.VotePlayer(payload.SenderName, payload.VotedName)

	if !ok {
		sender.SendStatusMessage(ws, server, 500, "Cannot vote for the player")
		return
	}

	sender.SendStatusMessage(ws, server, 200, "You have successfully voted")
	gameInstance.BroadcastState()
}

func HandleGameVoteWord(ws *websocket.Conn, server *api.WebSocketServer, message enums.GenericServerBoundMessage) {
	var payload enums.GameVoteWordMessage
	var err = json.Unmarshal(message.Payload, &payload)

	if err != nil {
		log.Println("Unmarshal error:", err)
		return
	}

	gameInstance := game.ManagerInstance.GetGame(payload.Code)
	if gameInstance == nil {
		sender.SendStatusMessage(ws, server, 404, "Game not found")
		return
	}

	ok := gameInstance.VoteWord(payload.Name, payload.Word)

	if !ok {
		sender.SendStatusMessage(ws, server, 500, "Cannot vote for the word")
		return
	}

	sender.SendStatusMessage(ws, server, 200, "You have successfully voted")
}

func HandleGamePlayerChatMessage(ws *websocket.Conn, server *api.WebSocketServer, message enums.GenericServerBoundMessage) {
	var payload enums.GamePlayerChatMessageMessage
	var err = json.Unmarshal(message.Payload, &payload)

	if err != nil {
		log.Println("Unmarshal error:", err)
		return
	}

	gameInstance := game.ManagerInstance.GetGame(payload.Code)
	if gameInstance == nil {
		sender.SendStatusMessage(ws, server, 404, "Game not found")
		return
	}

	gameInstance.AddChatMessage(payload.Name, payload.Message)

	gameInstance.BroadcastState()
}

func HandleGameCardPlacement(ws *websocket.Conn, server *api.WebSocketServer, message enums.GenericServerBoundMessage) {
	var payload enums.GameCardPlacementMessage
	var err = json.Unmarshal(message.Payload, &payload)

	if err != nil {
		log.Println("Unmarshal error:", err)
		return
	}

	gameInstance := game.ManagerInstance.GetGame(payload.Code)
	if gameInstance == nil {
		sender.SendStatusMessage(ws, server, 404, "Game not found")
		return
	}

	placed_cards := make([]game.PlacedCard, 0)
	for _, card := range payload.Placement.Cards {
		placed_cards = append(
			placed_cards,
			game.PlacedCard{
				Id:        card.Id,
				PositionX: card.PositionX,
				PositionY: card.PositionY,
				Rotation:  card.Rotation,
				Flipped:   card.Flipped,
			})
	}
	gameInstance.UpdateCardPlacement(placed_cards)
	gameInstance.StartVoteWord()
	gameInstance.BroadcastState()
}
