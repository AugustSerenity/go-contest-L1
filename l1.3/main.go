package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func worker(ch <-chan int, w int) {
	for val := range ch {
		fmt.Println(val, " worker number: ", w)
	}
}

func main() {
	n := flag.Int("n", 4, "количество воркеров")
	flag.Parse()

	ch := make(chan int)

	// Создаем N воркеров для чтения из главного канала
	for i := 0; i < *n; i++ {
		go worker(ch, i+1)
	}

	// Непрерывная запись данных в канал
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	// Ожидаем сигнала завершения программы
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt, syscall.SIGTERM)
	<-exitChan

	// Закрываем канал, чтобы завершить работу воркеров
	close(ch)
}
