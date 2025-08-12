package main

import "fmt"

type Human struct {
	name string
}

type Action struct {
	Human
}

func (h Human) Walk() {
	fmt.Printf("Today, %s walking\n", h.name)
}

func main() {
	man := &Action{Human: Human{"Nick"}}

	man.Walk()
}
