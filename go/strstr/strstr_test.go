package strstr

import "testing"

func Test_chunkString(t *testing.T) {
	input := "sadbutsadik"
	size := 3
	want := []string{"sad", "but", "sad", "ik"}

	got := chunkString(input, size)

	if len(got) != len(want) {
		t.Errorf("Expected %d chunks, got %d chunks", len(want), len(got))
	}
	for idx := range want {
		if got[idx] != want[idx] {
			t.Errorf("Chunk %d: expected %s, got %s", idx, want[idx], got[idx])
		}
	}
}

func Test_strStr(t *testing.T) {
	tests := map[string]struct {
		haystack string
		needle   string
		want     int
	}{
		"hello":        {haystack: "hello", needle: "ll", want: 2},
		"mississippi":  {haystack: "mississippi", needle: "issi", want: 1},
		"mississippi2": {haystack: "mississippi", needle: "issipi", want: -1},
		"leetcode":     {haystack: "leetcode", needle: "leeto", want: -1},
		"aaa":          {haystack: "aaa", needle: "aaaa", want: -1},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res := strStr(tc.haystack, tc.needle)
			if res != tc.want {
				t.Errorf("%s at index %d is not the first occurence of the needle", string(tc.haystack[res]), res)
			}
		})
	}
}
