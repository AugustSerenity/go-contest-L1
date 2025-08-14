package main

import "fmt"

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree", "dog"}

	m := make(map[string]bool)

	for i := 0; i < len(str); i++ {
		m[str[i]] = true
	}

	fmt.Println(m)

}
