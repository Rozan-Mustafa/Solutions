// Package clock contains the implementation of the Exercism's go exercise 'clock'.
package clock

import (
	"fmt"
	"strconv"
)

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

// The number of minutes of a day
const dayInMinutes = 60 * 24

type Clock struct {
	minute int
}

func New(hour, minute int) Clock {
	return Clock{0}.Add((hour * 60) + minute)
}

func (c Clock) String() string {
	hour := c.minute / 60
	for hour > 23 {
		hour -= 24
	}

	// process the minutes
	minute := c.minute % 60

	return fmt.Sprintf("%v:%v", addLeadingZero(hour), addLeadingZero(minute))
}

func (c Clock) Add(minutes int) Clock {
	// sum the minutes
	c.minute += minutes

	for c.minute < 0 {
		c.minute += dayInMinutes
	}

	// ignore additional days
	for c.minute > dayInMinutes {
		c.minute -= dayInMinutes
	}

	return c
}

func addLeadingZero(value int) string {
	result := strconv.Itoa(value)
	if value < 10 {
		result = "0" + result
	}
	return result
}
