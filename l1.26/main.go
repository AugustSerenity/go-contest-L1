package main

import (
	"fmt"
	"unicode"
)

func UniqCheck(s string) bool {
	m := make(map[rune]bool)

	for _, val := range s {
		lower := unicode.ToLower(val)
		if m[lower] {
			return false
		}
		m[lower] = true
	}

	return true
}

func main() {
	str := "1heL1o"

	fmt.Println(UniqCheck(str))
}
