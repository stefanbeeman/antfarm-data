package main

import (
	"github.com/stefanbeeman/antfarm"
)

func main() {
	world := antfarm.MakeWorld(".", 10, 10, 1)
	world.Run(10000)
}
