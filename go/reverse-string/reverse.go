package reverse

func String(s string) string {
	n := len(s)
	str := make([]rune, n)
	for _, rune := range s {
		n--
		str[n] = rune
	}
	return string(str[n:])

}
