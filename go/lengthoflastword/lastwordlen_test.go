package lengthoflastword

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	tests := map[string]struct {
		s    string
		want int
	}{
		"1": {"Hello World", 5},
		"2": {"   fly me   to   the moon  ", 4},
		"3": {"luffy is still joyboy", 6},
		"4": {"a", 1},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := lengthOfLastWord(tc.s)

			if got != tc.want {
				t.Errorf("got %d instead of %d", got, tc.want)
			}
		})
	}
}
