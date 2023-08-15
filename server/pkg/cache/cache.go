package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var Ctx = context.Background()
var Client *redis.Client

func Setup() {
	Client = redis.NewClient(&redis.Options{
		Addr:       os.Getenv("ADDR"),
		Password:   "",
		DB:         0,
		MaxRetries: 2,
	})

	log.Printf("%v", *Client.Ping(Ctx))
}
