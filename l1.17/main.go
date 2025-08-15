package main

import "fmt"

func BinarySearch(nums []int, key int) int {
	lo := 0
	hi := len(nums) - 1
	for lo <= hi {
		mi := lo + (hi-lo)/2
		if nums[mi] == key {
			return mi
		} else if nums[mi] < key {
			lo = mi + 1
		} else {
			hi = mi - 1
		}
	}

	return -1

}

func main() {
	nums := []int{1, 4, 5, 8, 9, 11}

	key := 5

	fmt.Println(BinarySearch(nums, key))
}
