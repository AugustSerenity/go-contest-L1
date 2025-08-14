package main

import (
	"fmt"
	"strings"
)

// Проблемы в исходном коде:
// 1. Утечка памяти: justString = v[:100] создает новую строку, но она все еще ссылается
//    на тот же базовый массив, что и v, поэтому огромная строка не может быть очищена GC
// 2. Возможная паника: если createHugeString вернет строку короче 100 символов,
//    операция v[:100] вызовет runtime error (slice bounds out of range)

// Исправленная реализация:
var justString string

func createHugeString(n int) string {
	return strings.Repeat("Go", n)
}

func someFunc() {
	v := createHugeString(2 << 10)
	// Создаем новый буфер нужного размера и копируем в него данные
	tmp := make([]byte, 100)
	// copy автоматически ограничит копирование длиной v
	copy(tmp, v[:100])
	// создаем новую строку, не связанную с v
	justString = string(tmp)
}

func main() {
	someFunc()
	fmt.Println(justString)
}
