package divide

import "math"

func divide(dividend int, divisor int) int {
	if dividend == divisor {
		return 1
	}
	if dividend == math.MinInt32 {

	}
	return -1
}

func divide1(dividend int, divisor int) int {
	if dividend > int(math.Pow(2, 31)-1) {
		return int(math.Pow(2, 31) - 1)
	}
	if dividend < -int(math.Pow(2, 31)) {
		return -int(math.Pow(2, 31))
	}
	n := 0
	sign := dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0
	dive := math.Abs(float64(dividend))
	divi := math.Abs(float64(divisor))
	for dive >= divi {
		n++
		dive -= divi
	}
	if sign {
		return n
	}
	return -n
}
