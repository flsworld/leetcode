package arraysign

import (
	"fmt"
	"slices"
)

func arraySign0(nums []int) int {
	product := 1.0
	for i := 0; i < len(nums); i++ {
		product *= float64(nums[i])
	}
	if product > 0 {
		return 1
	} else if product < 0 {
		return -1
	}
	return 0
}

func arraySign(nums []int) int {
	if slices.Contains(nums, 0) {
		return 0
	}
	k := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			k *= -1
		}
	}
	fmt.Println(k)
	return k
}
