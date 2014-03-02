package main

import (
	"github.com/stefanbeeman/antfarm"
)

func main() {
	game := af.MakeGame(".", 20, 20, 1)
	game.Start()
}
