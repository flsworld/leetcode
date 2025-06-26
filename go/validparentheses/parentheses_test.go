package validparentheses

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := map[string]struct {
		s    string
		want bool
	}{
		"1": {"()", true},
		"2": {"()[]{}", true},
		"3": {"(]", false},
		"4": {"([])", true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := isValid(tc.s)

			if got != tc.want {
				t.Errorf("got %t instead of %t", got, tc.want)
			}
		})
	}
}
