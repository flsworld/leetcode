package arraysign

import "testing"

func Test_arraySign(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want int
	}{
		"1": {[]int{-1, -2, -3, -4, 3, 2, 1}, 1},
		"2": {[]int{1, 5, 0, 2, -3}, 0},
		"3": {[]int{-1, 1, -1, 1, -1}, -1},
		"4": {[]int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24}, -1},
		"5": {[]int{-1, -2, -3, -4, 3, 2, 1}, 1},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := arraySign(tc.nums)

			if got != tc.want {
				t.Errorf("got %d instead of %d", got, tc.want)

			}
		})
	}
}
