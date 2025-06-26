package arithmeticprogression

import "testing"

func Test_canMakeArithmeticProgression(t *testing.T) {
	tests := map[string]struct {
		arr  []int
		want bool
	}{
		"1": {[]int{3, 5, 1}, true},
		"2": {[]int{1, 2, 4}, false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := canMakeArithmeticProgression(tc.arr)

			if got != tc.want {
				t.Errorf("got %t instead of %t", got, tc.want)

			}
		})
	}
}
