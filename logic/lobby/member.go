package lobby

import (
	"danny-backend/ws/dto/lobby"
	"github.com/gorilla/websocket"
)

type Member struct {
	Name       string
	Ready      bool
	Connection *websocket.Conn
}

func (m *Member) ToDto() lobby.MemberDto {
	return lobby.MemberDto{
		Name:  m.Name,
		Ready: m.Ready,
	}
}

func MemberFromDto(dto lobby.MemberDto, ws *websocket.Conn) Member {
	return Member{
		Name:       dto.Name,
		Ready:      dto.Ready,
		Connection: ws,
	}
}
