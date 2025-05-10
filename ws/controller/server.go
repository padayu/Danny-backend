package controller

import (
	"danny-backend/logic/game"
	"danny-backend/logic/lobby"
	"danny-backend/ws/api"
	"danny-backend/ws/enums"
	"danny-backend/ws/handler"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var handlerMap = map[enums.ServerBoundMessage]api.HandlerFunc{
	enums.RestoreConnection: handler.HandleRestoreConnection,

	enums.LobbyCreate:       handler.HandleLobbyCreate,
	enums.LobbyJoin:         handler.HandleLobbyJoin,
	enums.LobbyMemberUpdate: handler.HandleLobbyMemberUpdate,
	enums.LobbyKick:         handler.HandleLobbyKick,
	enums.LobbyLeave:        handler.HandleLobbyLeave,
	enums.LobbyStartGame:    handler.HandleLobbyStartGame,

	enums.GameVoteWord:          handler.HandleGameVoteWord,
	enums.GameVotePlayer:        handler.HandleGameVotePlayer,
	enums.GameCardPlacement:     handler.HandleGameCardPlacement,
	enums.GamePlayerChatMessage: handler.HandleGamePlayerChatMessage,
}

func NewWebSocketServer() *api.WebSocketServer {
	return &api.WebSocketServer{
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		Clients:         make(map[*websocket.Conn]bool),
		MessageHandlers: handlerMap,
		MessageSender:   sendGenericMessage,
	}
}

var Server *api.WebSocketServer

func Start() {
	Server = NewWebSocketServer()
	lobby.InitManager(Server)
	game.InitManager(Server)

	http.HandleFunc("/ws", handleConnections)

	log.Println("http Server started on :1337")
	err := http.ListenAndServe(":1337", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	server := Server

	ws, err := server.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	server.Clients[ws] = true

	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			log.Println("Read message error:", err)
			server.Clients[ws] = false
			return
		}

		fmt.Printf("Received message: %s\n", p)

		if messageType != websocket.TextMessage {
			return
		}

		var message enums.GenericServerBoundMessage
		err = json.Unmarshal(p, &message)
		if err != nil {
			log.Println("Unmarshal error:", err)
			return
		}

		handleGenericMessage(ws, message, server)
	}
}

func handleGenericMessage(ws *websocket.Conn, message enums.GenericServerBoundMessage, server *api.WebSocketServer) {
	if handlerFunc, ok := server.MessageHandlers[message.Type]; ok {
		handlerFunc(ws, server, message)
	} else {
		log.Printf("No handlerFunc for message type %s", message.Type)
	}
}

func sendGenericMessage(ws *websocket.Conn, message enums.GenericClientBoundMessage) {
	var data, err = json.Marshal(message)
	if err != nil {
		log.Println("Marshal error:", err)
		return
	}

	err = ws.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		log.Println("Write message error:", err)
		return
	}

	fmt.Printf("Sent message: %s\n", data)
}
