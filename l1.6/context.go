package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan int)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Stopping a goroutine")
				ch <- 1
				return
			default:
				fmt.Println("Wating...")
			}
		}
	}(ctx)

	time.Sleep(500 * time.Millisecond)
	cancel()

	<-ch
}
