package main

import "testing"

type TestCase struct {
	x, answer int
}

func TestSqrt(t *testing.T) {
	var TestCases = []TestCase{
		{0, 0},
		{1, 1},
		{3, 1},
		{4, 2},
		{6, 2},
		{8, 2},
		{9, 3},
		{256, 16},
		{266, 16},
		{410, 20},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := mySqrt(tc.x)
			if got != tc.answer {
				t.Logf("wrong answer in input %v\n", tc.x)
				t.Logf("want %d got %d\n", tc.answer, got)
				t.Fail()
			}
		})
	}
}
