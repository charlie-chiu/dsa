package main

import "testing"

type TestCase struct {
	input  string
	answer int
}

func TestCountBinarySubstrings(t *testing.T) {
	var TestCases = []TestCase{
		{"", 0},
		{"aab", 2},
		{"bbbbb", 1},
		{"abcabcbb", 3},
		{"pwwkew", 3},
		{"dvdf", 3},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := lengthOfLongestSubstring(tc.input)
			if got != tc.answer {
				t.Errorf("wrong answer in input %q, want %d got %d", tc.input, tc.answer, got)
			}
		})
	}
}
