package rmvelement

import "testing"

func Test_removeElement(t *testing.T) {
	tests := map[string]struct {
		nums  []int
		val   int
		wantK int
	}{
		"1": {[]int{3, 2, 2, 3}, 3, 2},
		"2": {[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := removeElement(tc.nums, tc.val)

			if got != tc.wantK {
				t.Errorf("got %d instead of %d", got, tc.wantK)
			}
		})
	}
}
