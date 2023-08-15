package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/reazn/draw/pkg/utils"
	"math/rand"
)

type MessageResponse struct {
	Event   string `json:"event"`
	Message string `json:"message"`
	Name    string `json:"name"`
}

func (r *Room) HandleMessage(c *websocket.Conn, content map[string]interface{}) {

	messageData := &MessageContent{}
	utils.FillStruct(content, messageData)

	for roomId, conns := range r.rooms {
		for _, conn := range conns {
			if conn.Conn == c {
				for _, otherConn := range conns {
					if otherConn.Conn != c {
						fmt.Println(roomId, messageData.Message, rand.Int())
						data, err := json.Marshal(MessageResponse{Event: "message", Message: messageData.Message, Name: conn.Name})
						err = otherConn.Conn.WriteMessage(websocket.TextMessage, data)
						if err != nil {
							return
						}
					}
				}
				break
			}
		}
	}
}
