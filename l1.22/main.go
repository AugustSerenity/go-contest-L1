package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b big.Int
	a.SetString("1048576", 10) // 2^20
	b.SetString("2097152", 10) // 2^21

	var sum, sub, mul, div big.Int

	sum.Add(&a, &b)
	fmt.Println("Sum:", sum.String())

	sub.Sub(&a, &b)
	fmt.Println("Subtraction:", sub.String())

	mul.Mul(&a, &b)
	fmt.Println("Multiplication:", mul.String())

	if b.Sign() == 0 {
		fmt.Println("Division by zero!")
	} else {
		div.Div(&a, &b)
		fmt.Println("Division:", div.String())
	}
}
