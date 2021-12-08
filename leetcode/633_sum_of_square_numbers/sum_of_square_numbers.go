package main

import (
	"math"
)

func judgeSquareSum(c int) bool {
	l := 0
	r := int(math.Sqrt(float64(c)))

	for l <= r {
		sum := l*l + r*r
		if sum > c {
			r--
		} else if sum < c {
			l++
		} else {
			return true
		}
	}

	return false
}
