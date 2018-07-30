package raindrops

import "strconv"

func Convert(number int) string {
	var string string = ""
	if number%3 == 0 {
		string += "Pling"
	}
	if number%5 == 0 {
		string += "Plang"
	}
	if number%7 == 0 {
		string += "Plong"
	}
	if string == "" {
		string = strconv.Itoa(number)
	}
	return string
}
