package ws

import (
	"fmt"
	"github.com/gofiber/contrib/websocket"
)

func (r *Room) HandleCreateRoom(c *websocket.Conn, content interface{}) {

	createRoomContent, ok := content.(CreateRoomContent)
	if !ok {
		err := c.WriteMessage(websocket.TextMessage, []byte("Invalid content payload."))
		if err != nil {
			return
		}
	}

	fmt.Println(createRoomContent)
	err := c.WriteMessage(websocket.TextMessage, []byte("Created!"))
	if err != nil {
		return
	}
}
