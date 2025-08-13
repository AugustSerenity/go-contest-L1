package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	dur := flag.Int("d", 1, "время завершения программы")
	flag.Parse()

	ch := make(chan int)

	// Непрерывная запись данных в канал
	go func() {
		for i := 0; ; i++ {
			ch <- i
			time.Sleep(20 * time.Millisecond)
		}
	}()

	// Читаем из канала
	go func() {
		for val := range ch {
			fmt.Println("---", val)
		}
	}()

	// Запускаем таймер до завершения программы
	time.Sleep(time.Duration(*dur) * time.Second)

	close(ch)

}
