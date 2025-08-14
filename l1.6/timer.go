package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		timer := time.NewTimer(1 * time.Second)

		for {
			select {
			case <-timer.C:
				fmt.Println("timer worked, goroutine stop")
				return
			default:
				fmt.Println("Wating...")
			}
		}
	}()

	wg.Wait()

}
