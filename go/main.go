package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	// 128 64 32 16 8 4 2 1
	// 1   1  1  1  1 1 1 1

	return ""
}

func nextGreatestLetter(letters []byte, target byte) byte {
	return 'c'
}

func nextGreatestLetter0(letters []byte, target byte) byte {
	if target >= letters[len(letters)-1] {
		return letters[0]
	}
	i := 0
	for letters[i] <= target {
		i++
	}
	return letters[i]
}

func isRobotBounded(instructions string) bool {
	return false
}

func main() {
	//res := twoSum([]int{2, 7, 11, 15}, 9)
	//res := isPalindrome(-121)
	//res := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	//res := removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	//res := strStr("hello", "ll")
	fmt.Println(1 << 32)
	fmt.Println(byte('a'), byte('z'), byte('A'))       // 97 122 65
	fmt.Printf("%c", byte(90))                         // Z
	fmt.Printf("%c\n", byte(91))                       // Z
	fmt.Printf("%s\n", " ")                            // Z
	fmt.Printf("choucroute %c choucroute\n", byte(32)) // Z

	lastIndex := make([]int, 128)
	s := "abcabcbb"
	currentChar := s[0]
	lastIndex[currentChar] = 1
	fmt.Println(currentChar, string(currentChar), lastIndex)
}
