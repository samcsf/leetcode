package lc754

import "math"

func ReachNumber(target int) int {
	target = int(math.Abs(float64(target)))
	k := 0
	for target > 0 {
		k++
		target -= k
	}
	if target%2 == 0 {
		return k
	} else {
		return k + 1 + k%2
	}
}
