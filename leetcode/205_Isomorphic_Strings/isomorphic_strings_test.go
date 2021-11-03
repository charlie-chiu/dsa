package main

import "testing"

type TestCase struct {
	s, t   string
	answer bool
}

func TestIsIsomorphic(t *testing.T) {
	var TestCases = []TestCase{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"billy", "bill", false},
		{"ab", "ba", true},
		{"badc", "baba", false},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := isIsomorphic(tc.s, tc.t)
			if got != tc.answer {
				t.Logf("wrong answser for input %q and %q", tc.s, tc.t)
				t.Logf("want %v got %v", tc.answer, got)
				t.Fail()
			}
		})
	}
}
