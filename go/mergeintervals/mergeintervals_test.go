package mergeintervals

import "testing"

func Test_merge(t *testing.T) {
	tests := map[string]struct {
		intervals [][]int
		want      [][]int
	}{
		"1": {[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		"2": {[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		"3": {[][]int{{1, 4}, {5, 8}, {6, 9}, {7, 10}}, [][]int{{1, 4}, {5, 10}}},
		"4": {[][]int{{1, 3}}, [][]int{{1, 3}}},
		"5": {[][]int{{1, 4}, {0, 4}}, [][]int{{0, 4}}},
		"6": {[][]int{{1, 4}, {0, 0}}, [][]int{{0, 0}, {1, 4}}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := merge(tc.intervals)

			for i := range tc.want {
				for j := range tc.want[i] {
					if got[i][j] != tc.want[i][j] {
						t.Errorf("got %d instead of %d", got[i][j], tc.want[i][j])
					}
				}
			}
		})
	}
}
