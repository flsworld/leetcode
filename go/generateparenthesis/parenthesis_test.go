package generateparenthesis

import (
	"slices"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := map[string]struct {
		n    int
		want []string
	}{
		"1": {3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		"2": {1, []string{"()"}},
		"3": {2, []string{"(())", "()()"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := generateParenthesis(tc.n)

			for _, combination := range got {
				ok := slices.Contains(tc.want, combination)
				if !ok {
					t.Errorf("%v not in %v", combination, tc.want)
				}
			}
		})
	}
}
