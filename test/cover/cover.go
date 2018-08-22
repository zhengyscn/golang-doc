package cover

import "math"

func Sum(i, j int) int {
	if i < 0 {
		i = int(math.Abs(float64(i)))
	}
	if j < 0 {
		j = int(math.Abs(float64(j)))
	}

	return i + j
}
