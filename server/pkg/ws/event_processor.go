package ws

import (
	"fmt"
	"github.com/gofiber/contrib/websocket"
)

func (r *Room) ProcessWSEvent(c *websocket.Conn, payload Payload) {
	switch payload.Event {
	case Ping:
		HandlePing(c)
	case CreateRoom:
		r.HandleCreateRoom(c, payload.Content)
	case JoinRoom:
		r.HandleJoinRoom(c, payload.Content)
	case LeaveRoom:
		r.HandleLeaveRoom(c, payload.Content)
	case Message:
		r.HandleMessage(c, payload.Content)
	default:
		fmt.Println("wrong event")
		err := c.WriteMessage(websocket.TextMessage, []byte("Not an event."))
		if err != nil {
			break
		}
	}
}
