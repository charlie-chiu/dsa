package main

import "testing"

type TestCase struct {
	input, answer int
}

func TestCountPrime(t *testing.T) {
	var TestCases = []TestCase{
		{10, 4},
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 1},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := countPrimes(tc.input)
			if got != tc.answer {
				t.Logf("wrong answer in input %d", tc.input)
				t.Logf("want %d got %d", tc.answer, got)
				t.Fail()
			}
		})
	}
}
