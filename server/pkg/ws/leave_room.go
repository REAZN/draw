package ws

import (
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/reazn/draw/pkg/utils"
)

func (r *Room) HandleLeaveRoom(c *websocket.Conn, content map[string]interface{}) {

	leaveData := &LeaveRoomContent{}
	utils.FillStruct(content, leaveData)

	if conns, ok := r.rooms[leaveData.RoomId]; ok {
		var updateC []ConnData
		for _, conn := range conns {
			if conn.Conn != c {
				updateC = append(updateC, conn)
			}
		}
		r.rooms[leaveData.RoomId] = updateC
	}
	fmt.Println(content)
}
