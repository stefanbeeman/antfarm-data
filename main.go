package main

import (
	"fmt"
	"github.com/stefanbeeman/antfarm"
)

func main() {
	fmt.Println("Starting...")
	game := antfarm.MakeGame(".", 20, 20, 1)
	game.StartShell()
}
