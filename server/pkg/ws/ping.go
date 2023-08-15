package ws

import (
	"github.com/gofiber/contrib/websocket"
)

func HandlePing(c *websocket.Conn) {

	err := c.WriteMessage(websocket.TextMessage, []byte("Pong!"))
	if err != nil {
		return
	}
}
