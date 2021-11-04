package main

import "testing"

type TestCase struct {
	haystack, needle string
	answer           int
}

func TestStrStr(t *testing.T) {
	var TestCases = []TestCase{
		{"a", "a", 0},
		{"a", "b", -1},
		{"aa", "a", 0},
		{"a", "aa", -1},
		{"abc", "c", 2},
		{"hello", "ll", 2},
		{"cat", "rat", -1},
		{"hello", "ll", 2},
		{"hello", "", 0},
		{"abcde", "bcdd", -1},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := strStr(tc.haystack, tc.needle)
			if got != tc.answer {
				t.Logf("wrong answer in haystack %q and needle %q", tc.haystack, tc.needle)
				t.Logf("want %d got %d", tc.answer, got)
				t.Fail()
			}
		})
	}
}
