package movezeroes

import "testing"

func Test_moveZeroes(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want []int
	}{
		"1": {[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		"2": {[]int{0}, []int{0}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			moveZeroes(tc.nums)

			for i := range tc.nums {
				if tc.nums[i] != tc.want[i] {
					t.Errorf("got %d instead of %d", tc.nums[i], tc.want[i])
				}

			}
		})
	}
}
