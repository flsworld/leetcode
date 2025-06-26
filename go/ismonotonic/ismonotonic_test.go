package ismonotonic

import "testing"

func Test_isMonotonic(t *testing.T) {
	tests := map[string]struct {
		arr  []int
		want bool
	}{
		"1": {[]int{1, 2, 2, 3}, true},
		"2": {[]int{6, 5, 4, 4}, true},
		"3": {[]int{1, 3, 2}, false},
		"4": {[]int{3, 3, 3, 2}, true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := isMonotonic(tc.arr)

			if got != tc.want {
				t.Errorf("got %t instead of %t", got, tc.want)

			}
		})
	}
}
