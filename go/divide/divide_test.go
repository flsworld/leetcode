package divide

import "testing"

func Test_divide(t *testing.T) {
	tests := map[string]struct {
		dividend int
		divisor  int
		want     int
	}{
		"1": {10, 3, 3},
		"2": {-10, -3, 3},
		"3": {7, -3, -2},
		"4": {-2147483648, -1, 2147483647},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := divide(tc.dividend, tc.divisor)

			if got != tc.want {
				t.Errorf("got %d instead of %d", got, tc.want)
			}
		})
	}
}
