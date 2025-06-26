package rmvelement

import "fmt"

func removeElement(nums []int, val int) int {
	k := 0
	for i := range nums {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
func removeElement0(nums []int, val int) int {
	j := len(nums) - 1
	for i := 0; i <= j; {
		if nums[i] == val {
			if nums[j] == val {
				j--
			}
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
		i++
	}
	fmt.Println(nums)
	return len(nums) - 1 - j
}
