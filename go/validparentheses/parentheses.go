package validparentheses

import "slices"

func isValid(s string) bool {
	stack := make([]rune, 0)
	m := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, char := range s {
		if val, ok := m[char]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == val {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}

func isValid0(s string) bool {
	stack := make([]rune, 0)
	m := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, char := range s {
		if slices.Contains([]rune{'(', '[', '{'}, char) {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != m[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
