package main

import "testing"

type TestCase struct {
	s, t, answer string
}

func TestMinWindow(t *testing.T) {
	var TestCases = []TestCase{
		{"a", "a", "a"},
		{"a", "AA", ""},
		{"ADOBECODEBANC", "ABC", "BANC"},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := minWindow(tc.s, tc.t)
			if got != tc.answer {
				t.Logf("wrong answer in input s: %q t: %q\n", tc.s, tc.t)
				t.Logf("want %q got %q\n", tc.answer, got)
				t.Fail()
			}
		})
	}
}
