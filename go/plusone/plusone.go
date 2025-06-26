package plusone

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	digits = append([]int{1}, digits...)
	return digits
}

func plusOne0(digits []int) []int {
	last := len(digits) - 1
	if digits[last] != 9 {
		return append(digits[:last], digits[last]+1)
	}
	for last >= 0 && digits[last] == 9 {
		last--
	}
	if last == -1 {
		res := []int{1}
		for range len(digits) {
			res = append(res, 0)
		}
		return res
	}
	res := digits[:last]
	res = append(res, digits[last]+1)
	for range len(digits) - last - 1 {
		res = append(res, 0)
	}
	return res

	// tant que digit == 9
	//	decr last
	// si last == 0
	//	mettre 1 plus len(digits) - 1 0
	// sinon
	// ajouter 1 à last puis finir par des 0

	return nil
}
