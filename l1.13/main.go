package main

import "fmt"

func main() {
	a := 1
	b := 22

	a = a + b
	b = a - b
	a = a - b

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
}
