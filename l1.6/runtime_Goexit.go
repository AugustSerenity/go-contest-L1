package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("defer goroutine")

		fmt.Println("starting goroutine")
		runtime.Goexit()
		fmt.Println("this will't be printed")
	}()

	fmt.Println("start main")
	wg.Wait()
	//runtime.Gosched() //позволяет горутине выполниться
	fmt.Println("end main")
}
