package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/contrib/websocket"
)

type ListMembersResponse struct {
	Event string                  `json:"event"`
	Data  ListMembersResponseData `json:"data"`
}

type ListMembersResponseData struct {
	Members []string `json:"members"`
}

func (r *Room) GetCurrentRoom(conn *websocket.Conn) string {
	for roomID, conns := range r.rooms {
		for _, c := range conns {
			if c.Conn == conn {
				return roomID
			}
		}
	}
	return ""
}

func (r *Room) HandleListMembers(c *websocket.Conn, content interface{}) {

	roomID := r.GetCurrentRoom(c)
	if roomID == "" {
		return
	}

	conns, ok := r.rooms[roomID]
	if !ok {
		return
	}

	members := make([]string, len(conns))
	for i, connData := range conns {
		members[i] = connData.Name
	}
	fmt.Printf("%+v\n", members)

	data, err := json.Marshal(ListMembersResponse{Event: "members",
		Data: ListMembersResponseData{
			Members: members,
		}})

	err = c.Conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return
	}
}
