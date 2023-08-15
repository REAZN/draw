package ws

//type User struct {
//	Active bool
//}

//func Test() {
//	app.Get("/ws/penis", websocket.New(func(c *websocket.Conn) {
//		var (
//			messageType int
//			message     []byte
//			err         error
//		)
//
//		for {
//			messageType, message, err = c.ReadMessage()
//
//			if err != nil {
//				err := c.Close()
//				if err != nil {
//					break
//				}
//				log.Println(err)
//				break
//			}
//
//			payload := Payload{}
//			err := json.Unmarshal(message, &payload)
//
//			if err != nil {
//				log.Println("error: ", err)
//				break
//			}
//
//		}
//	}))
//}

//func Auugh() {
//	app.Get("/ws/penis", websocket.New(func(c *websocket.Conn) {
//
//		var (
//			messageType int
//			message     []byte
//			err         error
//		)
//
//		for {
//			messageType, message, err = c.ReadMessage()
//
//			if err != nil {
//				err := c.Close()
//				if err != nil {
//					break
//				}
//				log.Println(err)
//				break
//			}
//
//			room := Payload{}
//			err := json.Unmarshal(message, &room)
//
//			if err != nil {
//				log.Println("error: ", err)
//				break
//			}
//
//			result, err := cache.Client.Exists(cache.Ctx, "room:"+room.RoomId).Result()
//
//			fmt.Printf("%+v", result)
//
//			if err != nil {
//				log.Println("error: ", err)
//			}
//
//			if result > 0 {
//				err = c.WriteMessage(messageType, []byte("Room already created"))
//				if err != nil {
//					log.Println("error: ", err)
//					break
//				}
//				return
//			}
//
//			//result, err = cache.Client.HSet(cache.Ctx, "room:"+room.RoomId, room.User, "atruadadade").Result()
//
//			//err = c.WriteMessage(messageType, message)
//
//			//if err != nil {
//			//	log.Println("error: ", err)
//			//	break
//			//}
//		}
//	}))
//}
