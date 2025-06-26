package romantoint

func romanToInt(s string) int {
	mapping := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := 0
	for i := range s {
		if i+1 < len(s) && mapping[rune(s[i])] < mapping[rune(s[i+1])] {
			sum -= mapping[rune(s[i])]
		} else {
			sum += mapping[rune(s[i])]
		}
	}
	return sum
}

func romanToInt1(s string) int {
	mapping := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	sum := 0
	last := ""
	for i := range s {
		if string(s[i]) == "V" && last == "I" {
			sum += 4 - 1
		} else if string(s[i]) == "X" && last == "I" {
			sum += 9 - 1
		} else if string(s[i]) == "L" && last == "X" {
			sum += 40 - 10
		} else if string(s[i]) == "C" && last == "X" {
			sum += 90 - 10
		} else if string(s[i]) == "D" && last == "C" {
			sum += 400 - 100
		} else if string(s[i]) == "M" && last == "C" {
			sum += 900 - 100
		} else {
			sum += mapping[string(s[i])]
		}
		last = string(s[i])
	}
	return sum
}

func romanToInt0(s string) int {
	sum := 0
	last := ""
	for i := range s {
		switch string(s[i]) {
		case "V":
			if last == "I" {
				sum += 4 - 1
			} else {
				sum += 5
			}
		case "X":
			if last == "I" {
				sum += 9 - 1
			} else {
				sum += 10
			}
		case "L":
			if last == "X" {
				sum += 40 - 10
			} else {
				sum += 50
			}
		case "C":
			if last == "X" {
				sum += 90 - 10
			} else {
				sum += 100
			}
		case "D":
			if last == "C" {
				sum += 400 - 100
			} else {
				sum += 500
			}
		case "M":
			if last == "C" {
				sum += 900 - 100
			} else {
				sum += 1000
			}
		default:
			sum += 1
		}

		last = string(s[i])
	}
	return sum
}
