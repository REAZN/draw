package ws

import (
	"encoding/json"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

type ConnData struct {
	Conn *websocket.Conn
	Name string
}

type Room struct {
	rooms map[string][]ConnData
}

func NewRoom() *Room {
	return &Room{
		rooms: make(map[string][]ConnData),
	}
}

type ErrorResponse struct {
	Type string            `json:"type"`
	Code string            `json:"code"`
	Data ErrorResponseData `json:"data"`
}

type ErrorResponseData struct {
	Message string `json:"message"`
}

var app *fiber.App

func Setup() {

	app = fiber.New()
	room := NewRoom()

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Use(logger.New(logger.Config{
		Format: "[[${ip}]:${port} ${status} - ${method} ${path}]\n\n",
	}))

	app.Get("/ws/penis", websocket.New(func(c *websocket.Conn) {
		for {
			var (
				message []byte
				err     error
			)

			_, message, err = c.ReadMessage()

			if err != nil {
				err := c.Close()
				if err != nil {
					break
				}
				log.Println(err)
				break
			}

			var payload Payload
			if err := json.Unmarshal(message, &payload); err != nil {
				log.Println("error: ", err)
				continue
			}

			room.ProcessWSEvent(c, payload)
		}
	}))

	Run()
}

func Run() {
	log.Fatal(app.Listen(":3330"))
}
