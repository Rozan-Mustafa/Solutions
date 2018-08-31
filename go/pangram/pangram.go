package pangram

import "unicode"

func IsPangram(i string) bool {
	cases := map[rune]bool{}

	for _, c := range i {
		if unicode.IsUpper(c) {
			c = unicode.ToLower(c)
		}
		if c < 'a' || c > 'z' {
			continue
		}
		cases[c] = true
	}
	return len(cases) == 26

}
