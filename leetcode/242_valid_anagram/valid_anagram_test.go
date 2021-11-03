package main

import "testing"

type TestCase struct {
	s, t   string
	answer bool
}

func TestIsAnagram(t *testing.T) {
	var TestCases = []TestCase{
		{"anagram", "nagaram", true},
		{"senior", "junior", false},
		{"go", "php", false},
		{"ab", "a", false},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := isAnagram(tc.s, tc.t)
			if got != tc.answer {
				t.Logf("wrong answser for input %q and %q", tc.s, tc.t)
				t.Logf("want %v got %v", tc.answer, got)
				t.Fail()
			}
		})
	}
}
