package rmvduplicates

import "fmt"

// Two-pointer Technique:
// Use index to place unique elements in their correct position.
// Traverse the array using i, comparing each element to the previous one.
// If the current element differs from the previous, copy it to the index position and increment index.
func removeDuplicates(nums []int) int {
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}
	fmt.Println(nums)
	return j
}

func removeDuplicatesII(nums []int) int {
	k := 2

	if len(nums) <= 2 {
		return len(nums)
	}

	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[k-2] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

func removeDuplicatesIISwag(nums []int) int {
	index, occurrence := 1, 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			occurrence++
		} else {
			occurrence = 1
		}

		if occurrence <= 2 {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

func removeDuplicatesIIFail(nums []int) int {
	buffer, cpt := 0, 0

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			if buffer > 0 {
				nums = append(append(nums[:i], nums[i+1:]...), -1)
				cpt++
				i--
			} else {
				buffer++
			}
		} else {
			buffer = 0
		}
	}

	return len(nums) - cpt
}
