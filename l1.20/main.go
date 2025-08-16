package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"

	words := strings.Fields(str)

	for i := 0; i < len(words); i++ {
		fmt.Printf("%s ", words[len(words)-1-i])
	}

}
