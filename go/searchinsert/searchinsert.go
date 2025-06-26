package searchinsert

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	if target < nums[left] {
		return 0
	}
	if target > nums[right] {
		return right + 1
	}

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func searchInsert0(nums []int, target int) int {
	isAfter := func(idx int) bool {
		return nums[idx] < target
	}

	left, right := 0, len(nums)-1
	if target < nums[left] {
		return 0
	}
	if target > nums[right] {
		return right + 1
	}

	for right-left > 1 {
		mid := (left + right) / 2
		if isAfter(mid) {
			left = mid
		} else {
			right = mid
		}
	}

	if nums[left] < target {
		return left + 1
	}
	return left
}

func searchInsert2(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}
