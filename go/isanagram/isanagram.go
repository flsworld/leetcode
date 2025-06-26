package isanagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	chars := make(map[rune]int)
	for i := range s {
		chars[rune(s[i])]++
		chars[rune(t[i])]--
	}
	for _, v := range chars {
		if v != 0 {
			return false
		}
	}
	return true
}

func isAnagram0(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	memS := make(map[string]int)
	memT := make(map[string]int)
	for i := 0; i < len(s); i++ {
		memS[string(s[i])]++
		memT[string(t[i])]++
	}
	for k := range memS {
		if memS[k] != memT[k] {
			return false
		}
	}
	return true
}
