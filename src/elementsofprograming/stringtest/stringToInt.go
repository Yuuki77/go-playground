package stringtest

import "strconv"

func stringToInt(x int) string {
	isNegative := false
	if x < 0 {
		isNegative = true
	}

	var s string = ""

	for x != 0 {
		a := x % 10
		s += "0" + strconv.Itoa(Abs(a))
		x /= 10
	}

	digit := ""
	if isNegative {
		digit = "-"
	}

	s += digit
	return s
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
