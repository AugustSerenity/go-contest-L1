package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scan(&i)

	sl := []int{1, 2, 3, 5, 6}

	copy(sl[i:], sl[i+1:])
	sl = sl[:len(sl)-1]

	fmt.Println(sl)

}
