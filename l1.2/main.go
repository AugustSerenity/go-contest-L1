package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}

	//создаем WaitGroup группу для корректной отработки программы
	wg := &sync.WaitGroup{}

	//проходим по каждому значению массива
	for val := range nums {
		//добавляем в счетчик группы +1 на каждой итерации
		wg.Add(1)
		// Запускаем горутину с выводом квадрата элемента массива
		go func(num int) {
			// После выполнения отправляется сигнал о завершении
			defer wg.Done()

			fmt.Println(num * num)
		}(nums[val])

	}

	//Ожидаем выполнения всех горутин
	wg.Wait()

}
