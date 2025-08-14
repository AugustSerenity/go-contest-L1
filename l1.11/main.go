package main

import "fmt"

func main() {
	m := make(map[int]bool)

	sl1 := []int{10, 20, 40}
	sl2 := []int{11, 22, 40, 21, 23}

	var res []int

	for i := 0; i < len(sl1); i++ {
		m[sl1[i]] = true
	}

	for i := 0; i < len(sl2); i++ {
		if m[sl2[i]] {
			res = append(res, sl2[i])
		}
	}

	fmt.Println("res: ", res)

}
