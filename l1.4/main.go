package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	//создаем новый контекст с функцией отмены этого контекста
	ctx, cancel := context.WithCancel(context.Background())

	//в отдельной горутине ожидаем сигналы ОС
	go func() {
		//создаем канал для сигналов
		exit := make(chan os.Signal, 1)
		//создаем обработчик сигналов, сигналы отправляются в канал exit
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		//ожидаем сигнала (блокирует выполнение, до тех пор пока не придет сигнал)
		//при получении сигнала - разблокируется
		<-exit
		//вызывает функцию отмены контекста
		//везде, где проверяется ctx.Done(), получают уведомление
		cancel()
	}()

	//создаем WaitGroup группу для корректной отработки программы
	wg := &sync.WaitGroup{}

	//горутина №1
	//добавляем в счетчик группы +1
	wg.Add(1)
	// Запускаем горутину
	go func() {
		//После выполнения отправляется сигнал о завершении
		defer wg.Done()
		//запускам бесконечный цикл с select
		//либо ждет отмены контекста
		//либо выводит сообщение каждые 2 секунды
		for {
			select {
			//срабатывает при отмене контекста (вызове cancel())
			case <-ctx.Done():
				fmt.Println("Break the loop goroutine № 1")
				return
			//срабатывает каждые 2 секунды
			case <-time.After(2 * time.Second):
				fmt.Println("Hello in a loop goroutine № 1")
			}

		}
	}()

	//аналогично горутине №1
	//горутина №2
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Break the loop goroutine № 2")
				return
			case <-time.After(2 * time.Second):
				fmt.Println("Hello in a loop goroutine № 2")
			}

		}
	}()

	wg.Wait()
	fmt.Println("All goroutines stopped")
}

//использую контекст т.к.:
//он безопасно передаётся в горутины
//позволяет централизованно отменить работу всех горутин (через ctx.Done())
//удобен для таймаутов, дедлайнов и отмены
//так же рекомендован стандартом Go для управления жизненным циклом процессов
