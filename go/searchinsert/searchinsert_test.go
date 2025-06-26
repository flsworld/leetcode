package searchinsert

import "testing"

func Test_searchInsert(t *testing.T) {
	tests := map[string]struct {
		nums   []int
		target int
		want   int
	}{
		"1": {[]int{1, 3, 5, 6}, 5, 2},
		"2": {[]int{1, 3, 5, 6}, 2, 1},
		"3": {[]int{1, 3, 5, 6}, 7, 4},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := searchInsert(tc.nums, tc.target)
			if got != tc.want {
				t.Errorf("got %d instead of %d", got, tc.want)
			}
		})
	}
}
