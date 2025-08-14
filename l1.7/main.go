package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	wg := &sync.WaitGroup{}

	table := make(map[int]int)

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()

			mu.Lock()
			table[x] = x
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Println(table)
}
