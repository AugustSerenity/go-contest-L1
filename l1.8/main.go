package main

import (
	"fmt"
)

func main() {
	var num, i, v int64

	fmt.Println("enter number:")
	fmt.Scanln(&num)
	fmt.Printf("%064b\n", num)

	fmt.Println("enter bit number:")
	fmt.Scanln(&i)

	fmt.Println("bit to set")
	fmt.Scanln(&v)

	switch v {
	case 0:
		num &= ^0 << i
	case 1:
		num |= 1 << i

	}
	fmt.Printf("%064b\n", num)
	fmt.Println(num)
}
