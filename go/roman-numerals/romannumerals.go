package romannumerals

import "errors"

func ToRomanNumeral(arabic int) (roman string, err error) {
	if arabic < 1 || arabic > 3000 {
		return "", errors.New("numbers lower than 1 and higher than 3000 can not be handled")
	} else {
		roman = toarabic(arabic)
	}
	return roman, nil
}
func toarabic(num int) string {

	if num >= 1000 {
		return "M" + toarabic(num-1000)
	}
	if num >= 900 {
		return "CM" + toarabic(num-900)
	}
	if num >= 500 {
		return "D" + toarabic(num-500)
	}
	if num >= 400 {
		return "CD" + toarabic(num-400)
	}
	if num >= 100 {
		return "C" + toarabic(num-100)
	}
	if num >= 90 {
		return "XC" + toarabic(num-90)
	}
	if num >= 50 {
		return "L" + toarabic(num-50)
	}
	if num >= 40 {
		return "XL" + toarabic(num-40)
	}
	if num >= 10 {
		return "X" + toarabic(num-10)
	}
	if num >= 9 {
		return "IX" + toarabic(num-9)
	}
	if num >= 5 {
		return "V" + toarabic(num-5)
	}
	if num >= 4 {
		return "IV" + toarabic(num-4)
	}
	if num >= 1 {
		return "I" + toarabic(num-1)
	}
	return ""
}
