package main

import "testing"

type TestCase struct {
	input  string
	answer int
}

func TestCountBinarySubstrings(t *testing.T) {
	var TestCases = []TestCase{
		{"a", 1},
		{"bb", 2},
		{"aabbb", 5},
		{"abccccdd", 7},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := longestPalindrome(tc.input)
			if got != tc.answer {
				t.Errorf("wrong answer in input %q, want %d got %d", tc.input, tc.answer, got)
			}
		})
	}
}
