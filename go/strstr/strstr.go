package strstr

func chunkString(s string, size int) []string {
	var res []string
	for i := 0; i < len(s); i += size {
		if i+size > len(s) {
			res = append(res, s[i:])
		} else {
			res = append(res, s[i:i+size])
		}
	}
	return res
}

// Find the Index of the First Occurrence in a String
func strStrNoob(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	for i := range haystack {
		if haystack[i] == needle[0] {
			if len(haystack)-i < len(needle) {
				return -1
			}
			ok := true
			for j := 0; j < len(needle) && i+j < len(haystack); j++ {
				if haystack[i+j] != needle[j] {
					ok = false
					break
				}
			}
			if ok {
				return i
			}
		}
	}
	return -1
}

func strStr(haystack string, needle string) int {
	for i := 0; i+len(needle) <= len(haystack); i++ {
		if haystack[i:len(needle)+i] == needle {
			return i
		}
	}
	return -1
}
