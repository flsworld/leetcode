package mergealternately

func mergeAlternately(word1 string, word2 string) string {
	res := ""
	n := min(len(word1), len(word2))
	for i := 0; i < n; i++ {
		res += string(word1[i]) + string(word2[i])
	}
	if len(word1) > len(word2) {
		res += word1[n:]
	} else {
		res += word2[n:]
	}
	return res
}
