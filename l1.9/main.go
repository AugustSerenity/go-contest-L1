package main

import (
	"fmt"
	"sync"
)

func main() {
	num := []int{1, 2, 3, 4, 5}

	ch1 := make(chan int)
	ch2 := make(chan int)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(num []int, ch1 chan int) {
		defer wg.Done()

		for _, val := range num {
			ch1 <- val
		}

		close(ch1)
	}(num, ch1)

	wg.Add(1)
	go func(ch1, ch2 chan int) {
		defer wg.Done()

		for val := range ch1 {
			ch2 <- val * 2
		}

		close(ch2)
	}(ch1, ch2)

	for v := range ch2 {
		fmt.Println(v)
	}

	wg.Wait()
}
