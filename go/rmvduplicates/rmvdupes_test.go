package rmvduplicates

import "testing"

func Test_removeDuplicatesII(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want int
	}{
		"1": {[]int{1, 1, 1, 2, 2, 3}, 5},
		"2": {[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7},
		"3": {[]int{0, 0, 0, 0, 0}, 2},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := removeDuplicatesII(tc.nums)

			if got != tc.want {
				t.Errorf("got %d instead of %d", got, tc.want)

			}
		})
	}
}
