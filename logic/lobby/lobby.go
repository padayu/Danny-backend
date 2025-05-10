package lobby

import (
	"danny-backend/ws/dto/lobby"
	"danny-backend/ws/sender"
	"github.com/gorilla/websocket"
)

type Lobby struct {
	Code    string
	Members []*Member
	Host    string
}

func (l *Lobby) ToDto() lobby.StateDto {
	members := make([]lobby.MemberDto, 0)
	for _, member := range l.Members {
		members = append(members, member.ToDto())
	}

	return lobby.StateDto{
		Code:    l.Code,
		Members: members,
		Host:    l.Host,
	}
}

func (l *Lobby) AddMember(name string, ws *websocket.Conn) bool {
	for _, m := range l.Members {
		if m.Name == name {
			return false
		}
	}

	l.Members = append(l.Members, &Member{
		Name:       name,
		Ready:      false,
		Connection: ws,
	})

	return true
}

func (l *Lobby) RemoveMember(name string) []*Member {
	newMembers := make([]*Member, 0)
	deletedMembers := make([]*Member, 0)
	for _, m := range l.Members {
		if m.Name != name {
			newMembers = append(newMembers, m)
		} else {
			deletedMembers = append(deletedMembers, m)
		}
	}

	l.Members = newMembers

	return deletedMembers
}

func (l *Lobby) UpdateMember(member Member) bool {
	for i, m := range l.Members {
		if m.Name == member.Name {
			*l.Members[i] = member
			return true
		}
	}

	return false
}

func (l *Lobby) SendState(member *Member) {
	sender.SendLobbyInfoMessage(member.Connection, ManagerInstance.Server, l.ToDto())
}

func (l *Lobby) BroadcastState() {
	for _, member := range l.Members {
		l.SendState(member)
	}
}

func (l *Lobby) BroadcastStart() {
	for _, member := range l.Members {
		sender.SendLobbyStartedMessage(member.Connection, ManagerInstance.Server)
	}
}

func (l *Lobby) RestoreConnection(name string, connection *websocket.Conn) bool {
	for _, m := range l.Members {
		if m.Name == name {
			m.Connection = connection
			l.SendState(m)
			return true
		}
	}

	return false
}
