package main

import "testing"

type TestCase struct {
	s          string
	dictionary []string
	answer     string
}

func TestFindLongestWord(t *testing.T) {
	var TestCases = []TestCase{
		{"hello", []string{"a", "b", "c"}, ""},
		{"abce", []string{"abe", "abc"}, "abc"},
		{"hello", []string{"ate", "bard", "cat", "bat"}, ""},
		{"abpcplea", []string{"a", "b", "c"}, "a"},
		{"abpcplea", []string{"ale", "apple", "monkey", "plea"}, "apple"},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := findLongestWord(tc.s, tc.dictionary)
			if got != tc.answer {
				t.Logf("wrong answer in input %q, %v\n", tc.s, tc.dictionary)
				t.Logf("want %q got %q\n", tc.answer, got)
				t.Fail()
			}
		})
	}
}
