package utils

import (
	"math/rand"
	"time"
)

func generateRoomId() string {
	length := 5
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	charset := "ABCDEFGHJKLPQRSTUVWXYZ23456789"

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seed.Intn(len(charset))]
	}

	return string(b)
}
