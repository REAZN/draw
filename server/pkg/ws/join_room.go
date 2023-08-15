package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/reazn/draw/pkg/utils"
)

func (r *Room) HandleJoinRoom(c *websocket.Conn, content map[string]interface{}) {

	fmt.Println(content)
	joinData := &JoinRoomContent{}
	utils.FillStruct(content, joinData)

	fmt.Println(*joinData)

	//Prevent the same connection from joining the same room twice
	if conns, ok := r.rooms[joinData.RoomId]; ok {
		for _, conn := range conns {
			if conn.Conn == c {
				data, err := json.Marshal(ErrorResponse{Type: "error", Code: "418",
					Data: ErrorResponseData{
						Message: "You have already joined this room.",
					}},
				)
				err = c.WriteMessage(websocket.TextMessage, data)
				if err != nil {
					return
				}
				return
			}
		}
	}

	r.rooms[joinData.RoomId] = append(r.rooms[joinData.RoomId], ConnData{Conn: c, Name: joinData.Name})
	fmt.Printf("Joined room %+v\n", r.rooms)

	var name []string
	if conns, ok := r.rooms[joinData.RoomId]; ok {
		name = make([]string, len(conns))
		for i, conn := range conns {
			name[i] = conn.Name
		}
	}
	connections, err := json.Marshal(name)
	err = c.WriteMessage(websocket.TextMessage, connections)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}
