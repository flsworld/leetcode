package mergeintervals

import (
	"fmt"
	"slices"
	"sort"
)

// Sort before working on the slice
// sort.Slice allows custom comparison function while slices.Sort doesn't
// Use index based on the length of manipulated slice instead of indices of the loop because
// it can lead to an Out Of Range

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] ||
			intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1]
	})

	res := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		if len(res) == 0 || intervals[i][0] > res[len(res)-1][1] {
			// no overlapping
			res = append(res, intervals[i])
		} else {
			// overlapping
			segment := []int{res[len(res)-1][0], max(res[len(res)-1][1], intervals[i][1])}
			res[len(res)-1] = segment
		}
		fmt.Println(res)
	}
	fmt.Println(res)
	return res
}

func mergeFirstTry(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}

	ref := intervals[0]
	var output [][]int

	for i := 1; i < len(intervals); i++ {
		switch {
		case intervals[i][0] <= ref[0]:
			if intervals[i][1] > ref[1] || intervals[i][1] < ref[0] {
				ref = []int{intervals[i][0], intervals[i][1]}
			} else {
				ref = []int{intervals[i][0], ref[1]}
			}
		case ref[0] < intervals[i][0] && intervals[i][0] <= ref[1]:
			if intervals[i][1] > ref[1] {
				ref = []int{ref[0], intervals[i][1]}
			}
		default:
			if len(output) == 0 {
				output = append(output, ref)
			}
			ref = intervals[i]
		}

		output = append(output, ref)
	}
	fmt.Println(output)

	out := [][]int{output[len(output)-1]}
	for i := len(output) - 2; i > -1; i-- {
		if output[i][0] != output[i+1][0] {
			out = append(out, output[i])
		}
	}
	fmt.Println(out)
	slices.Reverse(out)
	return out
}
