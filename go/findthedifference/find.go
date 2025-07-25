package findthedifference

import "fmt"

// You are given two strings s and t.
//
// String t is generated by random shuffling string s and then add one more letter at a random
// position.
//
// Return the letter that was added to t.

func findTheDifference0(s string, t string) byte {
	if len(s) == 0 {
		return t[0]
	}

	ms := make(map[rune]int)
	mt := make(map[rune]int)
	m := make(map[rune]int)

	for i := 0; i < len(s); i++ {
		ms[rune(s[i])]++
		mt[rune(t[i])]++
	}

	for k := range mt {
		m[k] = mt[k] - ms[k]
	}
	fmt.Println(m)
	for k, v := range m {
		if v == 1 {
			return byte(k)
		}
	}
	return t[len(t)-1]
}

func findTheDifference(s string, t string) byte {
	n := len(s)
	totalBytesS := t[0]
	for i := 1; i < n+1; i++ {
		totalBytesS += t[i]
	}
	fmt.Println(totalBytesS)

	for j := 0; j < n; j++ {
		totalBytesS -= s[j]
	}
	fmt.Println(totalBytesS)

	return totalBytesS
}
