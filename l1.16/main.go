package main

import "fmt"

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	base := nums[len(nums)/2]
	var less, more []int

	for i := range nums {
		if nums[i] < base {
			less = append(less, nums[i])
		} else if nums[i] > base {
			more = append(more, nums[i])
		}
	}

	less = quickSort(less)
	more = quickSort(more)

	less = append(less, base)

	return append(less, more...)

}

func main() {
	nums := []int{8, 3, 2, 4, 67, 33, 34, 11, 6}
	fmt.Println(nums)

	nums = quickSort(nums)
	fmt.Println(nums)
}
