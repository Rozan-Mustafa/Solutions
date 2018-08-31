package grains

import "errors"

func Square(input int) (grains uint64, invalid error) {
	if input <= 0 || input > 64 {
		invalid = errors.New("Invalid input")
	} else {
    for 
		grains = 1 << uint(input-1)
	}
	return
}

func Total() uint64 {
	return 1<<64 - 1
}
