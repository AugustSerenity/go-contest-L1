package main

import (
	"fmt"
	"time"
)

func main() {
	var x bool

	go func() {
		for {
			fmt.Println("starting goroutine")

			if x {
				fmt.Println("stop goroutine")
				return
			}
		}
	}()

	time.Sleep(500 * time.Millisecond)
	x = true
	time.Sleep(500 * time.Millisecond)

}
