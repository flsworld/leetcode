package ismonotonic

func isMonotonic(nums []int) bool {
	sign := 0
	for j := 1; j < len(nums); j++ {
		diff := nums[j] - nums[j-1]
		switch {
		case diff == 0:
			continue
		case diff < 0 && sign == 0:
			sign = -1
		case diff > 0 && sign == 0:
			sign = 1
		case diff < 0 && sign > 0 || diff > 0 && sign < 0:
			return false
		}
	}
	return true
}
