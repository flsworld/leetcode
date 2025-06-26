package search

import "testing"

func Test_search(t *testing.T) {
	tests := map[string]struct {
		nums   []int
		target int
		want   int
	}{
		"1": {[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		"2": {[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		"3": {[]int{2, 5}, 5, 1},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := search(tc.nums, tc.target)

			if got != tc.want {
				t.Errorf("got %d instead of %d", got, tc.want)

			}
		})
	}
}
