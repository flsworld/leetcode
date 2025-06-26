package twosum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if val, ok := m[diff]; ok {
			return []int{val, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}
