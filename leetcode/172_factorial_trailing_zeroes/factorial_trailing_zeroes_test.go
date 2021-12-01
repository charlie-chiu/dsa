package main

import "testing"

type TestCase struct {
	input, answer int
}

func TestTrailingZeroes(t *testing.T) {
	var TestCases = []TestCase{
		{3, 0},
		{5, 1},
		{16, 3},
		{20, 4},
		{25, 6},
		{30, 7},
		{129, 31},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := trailingZeroes(tc.input)
			if got != tc.answer {
				t.Logf("wrong answer in input %d\n", tc.input)
				t.Logf("want %d got %d\n", tc.answer, got)
				t.Fail()
			}
		})
	}
}
