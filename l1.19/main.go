package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "главрыба"

	var res string
	r := []rune(str)

	fullLen := utf8.RuneCountInString(str)

	for i := 0; i < fullLen; i++ {
		res += string(r[(fullLen-1)-i])
	}

	fmt.Println(res)
}
