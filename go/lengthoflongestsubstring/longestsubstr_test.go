package lengthoflongestsubstring

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := map[string]struct {
		s    string
		want int
	}{
		"1": {"abcabcbb", 3},
		"2": {"bbbbb", 1},
		"3": {"pwwkew", 3},
		"4": {"", 0},
		"5": {" ", 1},
		"6": {"au", 2},
		"7": {"aab", 2},
		"8": {"dvdf", 3},
		"9": {"anviaj", 5},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tc.s)

			if got != tc.want {
				t.Errorf("got %d instead of %d", got, tc.want)
			}

		})
	}
}
