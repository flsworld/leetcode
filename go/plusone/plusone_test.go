package plusone

import "testing"

func Test_plusOne(t *testing.T) {
	tests := map[string]struct {
		digits []int
		want   []int
	}{
		"1": {[]int{1, 2, 3}, []int{1, 2, 4}},
		"2": {[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		"3": {[]int{9}, []int{1, 0}},
		"4": {[]int{9, 9, 9}, []int{1, 0, 0, 0}},
		"5": {[]int{4, 8, 9, 9}, []int{4, 9, 0, 0}},
		"6": {[]int{8, 9, 9, 9}, []int{9, 0, 0, 0}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := plusOne(tc.digits)

			if len(got) != len(tc.want) {
				t.Errorf("got %d instead of %d", len(got), len(tc.want))
			}

			for i := range got {
				if got[i] != tc.want[i] {
					t.Errorf("got %d instead of %d", got[i], tc.want[i])
				}
			}
		})
	}
}
