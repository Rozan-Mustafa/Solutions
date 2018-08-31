package lsproduct

import (
	"errors"
	"strconv"
	"unicode"
)

func LargestSeriesProduct(digits string, limit int) (product int64, err error) {
	switch {
	case limit == 0:
		return 1, nil
	case limit < 0:
		return 0, errors.New("limit must be > 0")
	case limit > len(digits):
		return 0, errors.New("limit must be shorter than the lenght of digits")
	case !allNumbers(digits):
		return -1, errors.New("the digit string must contain only numbers")
	default:
		return largestProduct(digits, limit), nil
	}
}

func largestProduct(digits string, limit int) int64 {
	var max int64
	var product int64
	length := len(digits)
	for i := 0; i <= length-limit; i++ {
		substr := digits[i : limit+i]
		product = 1
		for _, d := range substr {
			multiplier, _ := strconv.Atoi(string(d))
			product *= int64(multiplier)
		}
		if product > max {
			max = product
		}
	}
	return max
}

func allNumbers(digits string) bool {
	for _, d := range digits {
		if !unicode.IsNumber(d) {
			return false
		}
	}
	return true
}
