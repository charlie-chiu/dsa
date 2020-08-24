package sort

import (
	"math"
)

// slower than first bubble sort, doesn't know why
// bubble up smallest (lightest) value to the top of list
func bubble2(input []int) {
	// slower than bubble, doesn't know why
	var length = len(input)
	for i := 0; i < length; i++ {
		for j := length - 1; j > i; j-- {
			prev := j - 1
			if input[j] < input[prev] {
				input[j], input[prev] = input[prev], input[j]
			}
		}
	}
}

func bubble(input []int) {
	var length = len(input)
	var sorted = 0
	var compareL int
	var compareR int

	for sorted < length {
		compareL = length - 2
		compareR = length - 1

		for compareL >= sorted {
			if input[compareL] > input[compareR] {
				input[compareL], input[compareR] = input[compareR], input[compareL]
			}

			compareL--
			compareR--
		}

		sorted++
	}
}

//search smallest value, swap with the leftmost value
func selection(input []int) {
	var leftmost int
	var minIndex int
	var minValue int

	for leftmost < len(input) {
		minIndex = leftmost
		minValue = math.MaxInt64
		for i := leftmost; i < len(input); i++ {
			if input[i] < minValue {
				minIndex = i
				minValue = input[i]
			}

		}
		input[leftmost], input[minIndex] = input[minIndex], input[leftmost]

		leftmost++
	}
}
