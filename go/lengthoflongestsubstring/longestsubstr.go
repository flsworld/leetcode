package lengthoflongestsubstring

import (
	"slices"
)

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	start := 0
	lastIndex := make([]int, 128)

	for end := 0; end < len(s); end++ {
		currentChar := s[end]
		if lastIndex[currentChar] > start {
			start = lastIndex[currentChar]
		}
		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
		lastIndex[currentChar] = end + 1
	}
	return maxLength
}
func lengthOfLongestSubstring0(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	m := make(map[rune]int)
	res := make([]int, 0)
	noRepetition := true
	for i := 0; i < len(s); {
		m[rune(s[i])]++
		if m[rune(s[i])] > 1 {
			if len(m) == 1 {
				m = make(map[rune]int)
				continue
			}
			res = append(res, len(m))
			m = make(map[rune]int)
			noRepetition = false
			i--
		} else {
			i++
		}
	}
	if noRepetition {
		return len(m)
	}
	return max(slices.Max(res), len(m))
}
