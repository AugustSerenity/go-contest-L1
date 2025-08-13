package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func(ch <-chan int) {
		for {
			select {
			case <-ch:
				fmt.Println("Stopping a goroutine")
				return
			default:
				fmt.Println("Wating...")
			}
		}
	}(ch)

	time.Sleep(500 * time.Millisecond)
	ch <- 1
}
