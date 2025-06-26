package romantoint

import "testing"

func Test_romanToInt(t *testing.T) {
	tests := map[string]struct {
		s    string
		want int
	}{
		"1": {"III", 3},
		"2": {"LVIII", 58},
		"3": {"MCMXCIV", 1994},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := romanToInt(tc.s)

			if got != tc.want {
				t.Errorf("got %d instead of %d", got, tc.want)

			}
		})
	}
}
