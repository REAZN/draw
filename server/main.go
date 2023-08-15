package main

import (
	"github.com/reazn/draw/pkg/cache"
	"github.com/reazn/draw/pkg/ws"
)

func main() {
	cache.Setup()
	ws.Setup()
}
