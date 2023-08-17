package ws

type Event string

const (
	Ping        Event = "ping"
	CreateRoom  Event = "create"
	JoinRoom    Event = "join"
	LeaveRoom   Event = "leave"
	ListMembers Event = "members"
	Message     Event = "message"
)

type Payload struct {
	Event   Event                  `json:"event"`
	Content map[string]interface{} `json:"data"`
}

type CreateRoomContent struct {
	Name string `json:"name"`
}

type JoinRoomContent struct {
	Name   string `json:"name"`
	RoomId string `json:"roomid"`
}

type LeaveRoomContent struct {
	RoomId string `json:"roomid"`
}

type MessageContent struct {
	Message string `json:"message"`
}
