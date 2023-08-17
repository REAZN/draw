package ws

import (
	"encoding/json"
	"github.com/gofiber/contrib/websocket"
)

type PingResponse struct {
	Event string           `json:"event"`
	Data  PingResponseData `json:"data"`
}

type PingResponseData struct {
	Message string `json:"message"`
}

func HandlePing(c *websocket.Conn) {

	data, err := json.Marshal(PingResponse{Event: "ping", Data: PingResponseData{Message: "Pong!"}})
	err = c.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return
	}
}
