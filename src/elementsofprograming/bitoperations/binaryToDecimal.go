package elementsofprograming

import (
	"math"
	"strconv"
)

func binaryToDecimal(x int) int {
	sum := 0
	s := strconv.Itoa(x)

	for index := 0; index < len(s); index++ {
		current, _ := strconv.Atoi(string(s[len(s)-1-index]))
		if current == 1 {
			sum += int(math.Pow(2, float64(index)))
		}
	}

	return sum
}
