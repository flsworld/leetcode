package lengthoflastword

func lengthOfLastWord(s string) int {
	i := len(s)
	res := 0
	for i > 0 {
		i--
		if string(s[i]) != " " {
			res++
		} else if res > 0 {
			return res
		}
	}
	return res
}
