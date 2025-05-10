package api

import (
	"danny-backend/ws/enums"
	"github.com/gorilla/websocket"
)

type HandlerFunc func(*websocket.Conn, *WebSocketServer, enums.GenericServerBoundMessage)
type SenderFunc func(*websocket.Conn, enums.GenericClientBoundMessage)

type WebSocketServer struct {
	Upgrader        websocket.Upgrader
	Clients         map[*websocket.Conn]bool
	MessageHandlers map[enums.ServerBoundMessage]HandlerFunc
	MessageSender   SenderFunc
}
