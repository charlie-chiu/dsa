package main

import "testing"

type TestCase struct {
	input  string
	output int
}

func TestCountSubStrings(t *testing.T) {
	var TestCases = []TestCase{
		{"a", 1},
		{"abc", 3},
		{"aaa", 6},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := countSubstrings(tc.input)
			if got != tc.output {
				t.Errorf("wrong answer in input %q, want %d got %d", tc.input, tc.output, got)
			}
		})
	}
}
