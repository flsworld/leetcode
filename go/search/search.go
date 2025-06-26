package search

func search(nums []int, target int) int {
	n := len(nums) - 1
	if nums[0] == target {
		return 0
	}
	if nums[n] == target {
		return n
	}
	left, right := 0, n
	for i := 0; i < (n/2)+1; i++ {
		mid := (left + right) / 2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			left = mid
		case nums[mid] > target:
			right = mid
		}
	}

	return -1
}
