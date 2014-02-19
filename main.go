package main

import (
	"github.com/stefanbeeman/antfarm"
)

type Stack interface {
	foo()
}

type Mack interface {
	Stack
	goo()
}

type StackMack struct{}

func main() {
	w := af.MakeWorld(".", 20, 20, 1)
	w.Start()
}
