package main

import "fmt"

func main() {
	var val interface{}
	val = 123124

	switch mytype := val.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan any, chan int, chan string, chan bool:
		fmt.Println("chan")

	default:
		fmt.Printf("%T", mytype)
	}
}
