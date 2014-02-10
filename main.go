package main

import (
	"antfarm"
	"fmt"
)

func main() {
	world := antfarm.MakeWorld(100, 100)
	world.Show()
	world.Run(1000)
	fmt.Println("done with simulation")
}
