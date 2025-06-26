package mergealternately

func mergeAlternatelyOpti(word1 string, word2 string) string {
	result := make([]byte, 0, len(word1)+len(word2))
	minLength := len(word1)
	if len(word2) < minLength {
		minLength = len(word2)
	}

	for i := 0; i < minLength; i++ {
		result = append(result, word1[i], word2[i])
	}

	if len(word1) > minLength {
		result = append(result, word1[minLength:]...)
	} else if len(word2) > minLength {
		result = append(result, word2[minLength:]...)
	}

	return string(result)
}

func mergeAlternately(word1 string, word2 string) string {
	res := ""
	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			res += string(word1[i])
		}
		if i < len(word2) {
			res += string(word2[i])
		}
	}
	return res
}

func mergeAlternately0(word1 string, word2 string) string {
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
