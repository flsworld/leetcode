package isanagram

import "testing"

func Test_isAnagram(t *testing.T) {
	tests := map[string]struct {
		s    string
		t    string
		want bool
	}{
		"1": {"anagram", "nagaram", true},
		"2": {"rat", "car", false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := isAnagram(tc.s, tc.t)

			if got != tc.want {
				t.Errorf("got %t instead of %t", got, tc.want)
			}
		})
	}
}
