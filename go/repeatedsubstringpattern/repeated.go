package repeatedsubstringpattern

import "fmt"

func repeatedSubstringPattern(s string) bool {
	_ = s + s
	return true
}

func repeatedSubstringPattern0(s string) bool {
	step := 1
	i := 1
	var pattern, mem string
	for step <= len(s)/2 {
		pattern = s[(i-1)*step : (i-1)*step+step]
		mem = s[i*step : i*step+step]
		if i*step+step == len(s) && mem == pattern {
			return true
		} else if mem == pattern {
			i++
			if i*step+step > len(s) {
				return false
			}
		} else {
			i = 1
			step++
			fmt.Println("step", step, "i", i)
		}
	}
	return false
}
