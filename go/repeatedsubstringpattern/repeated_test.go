package repeatedsubstringpattern

import "testing"

func Test_repeatedSubstringPattern(t *testing.T) {
	tests := map[string]struct {
		s    string
		want bool
	}{
		"1": {"abab", true},
		"2": {"aba", false},
		"3": {"abcabcabcabc", true},
		"4": {"aaa", true},
		"5": {"abaababaab", true},
		"6": {"aabaaba", false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := repeatedSubstringPattern(tc.s)

			if got != tc.want {
				t.Errorf("got %t instead of %t", got, tc.want)
			}
		})
	}
}
