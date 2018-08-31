// Package isogram contains the implementation of the Exercism's go exercise 'Isogram'
package isogram

import (
	s "strings"
)

// IsIsogram determines whether the given input is an isogram or not
func IsIsogram(str string) bool {
	str = s.ToUpper(str)
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] && str[i] > 64 && str[i] < 91 {
				return false
			}
		}
	}
	return true
}
