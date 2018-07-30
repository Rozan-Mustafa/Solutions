package scrabble

import s "strings"

func Score(str string) int {
	str = s.ToUpper(str)
	var points = []int{
		1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4,
		8, 4, 10}
	score := 0
	for i := 0; i < len(str); i++ {
		if str[i] == 32 {
			score += 0
		} else {
			score += points[str[i]-65]
		}
	}
	return score

}
