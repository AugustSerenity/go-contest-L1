package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []float64{-25.4, -27.0, 13.0, 19.0,
		15.5, 24.5, -21.0, 32.5}

	res := make(map[int][]float64)

	for _, temp := range nums {
		var key float64
		if temp > 0 {
			key = math.Floor(temp/10) * 10
		} else {
			key = math.Ceil(temp/10) * 10
		}
		res[int(key)] = append(res[int(key)], temp)
	}

	fmt.Println(res)
}
