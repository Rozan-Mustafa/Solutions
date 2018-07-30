package hamming

import "errors"

func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, errors.New("strings are not equal")
	}
	var Dist = 0
	for i := range a {
		if a[i] != b[i] {
			Dist = Dist + 1
		}
	}
	return Dist, nil
}
