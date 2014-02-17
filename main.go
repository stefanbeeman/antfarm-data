package main

import (
	"github.com/stefanbeeman/antfarm"
)

func main() {
	w := antfarm.MakeWorld(".", 20, 20, 1)
	w.Start()
}
