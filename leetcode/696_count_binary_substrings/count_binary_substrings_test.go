package main

import "testing"

type TestCase struct {
	input  string
	answer int
}

func TestCountBinarySubstrings(t *testing.T) {
	var TestCases = []TestCase{
		{"", 0},
		{"0", 0},
		{"1", 0},
		{"11", 0},
		{"10", 1},
		{"10101", 4},
		{"00110011", 6},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := countBinarySubstrings(tc.input)
			if got != tc.answer {
				t.Errorf("wrong answer in input %q, want %d got %d", tc.input, tc.answer, got)
			}
		})
	}
}
