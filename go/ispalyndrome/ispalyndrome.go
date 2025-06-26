package ispalyndrome

import (
	"fmt"
	"strconv"
)

func isPalindromeNoob(x int) bool {
	if x < 0 {
		return false
	}
	input := strconv.Itoa(x)
	if len(input)%2 == 0 {
		return false
	}
	pivot := len(input) / 2
	fmt.Println(pivot, len(input))
	for i := 0; i < pivot; i++ {
		if input[pivot+i] == input[pivot-i] {
			continue
		} else {
			return false
		}
	}
	return true
}

func isPalindrome2(x int) bool {
	xStr := strconv.Itoa(x)
	for idx, jdx := 0, len(xStr)-1; idx <= jdx; {
		if xStr[idx] != xStr[jdx] {
			return false
		}
		idx++
		jdx--
	}
	return true
}

func isPalindrome(x int) bool {
	reverseX := 0
	for num := x; num > 0; {
		rem := num % 10
		reverseX = reverseX*10 + rem
		num = num / 10
	}
	if x == reverseX {
		return true
	}
	return false
}
