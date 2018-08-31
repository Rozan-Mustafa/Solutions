package luhn

import (
	"strconv"
	s "strings"
)

func Valid(str string) bool {
	var Digit int
	var nextDigit int
	str = s.Replace(str, " ", "", -1)
	sum := 0
	if !isAllowed(str) {
		return false
	}
	for i := len(str) - 1; i > -1; i-- {
		Digit, _ = strconv.Atoi(string(str[i]))
		sum += Digit
		i = i - 1
		if i >= 0 {
			nextDigit, _ = strconv.Atoi(string(str[i]))
			sum += doubleDigit(nextDigit)
		}
	}
	return sum%10 == 0
}

func isAllowed(str string) bool {
	if _, err := strconv.Atoi(str); err == nil {
		if len(str) > 1 {
			return true
		}
	}
	return false
}
func doubleDigit(input int) int {
	result := input * 2
	if result > 9 {
		result -= 9
	}
	return result
}
