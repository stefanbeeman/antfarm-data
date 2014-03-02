package main

import (
	"github.com/stefanbeeman/antfarm"
)

func main() {
	game := antfarm.MakeGame(".", 20, 20, 1)
	game.Start()
}
