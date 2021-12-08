package main

import (
	"fmt"
	"testing"
)

type TestCase struct {
	s      string
	answer bool
}

func TestValidPalindrome(t *testing.T) {
	var TestCases = []TestCase{
		{"aa", true},
		{"aba", true},
		{"abc", false},
		{"bc", true},
		{"abca", true},
		{"caba", true},
	}

	for _, tc := range TestCases {
		t.Run(fmt.Sprintf("input %q", tc.s), func(t *testing.T) {
			got := validPalindrome(tc.s)
			if got != tc.answer {
				t.Logf("wrong answer in input %q\n", tc.s)
				t.Logf("want %v got %v\n", tc.answer, got)
				t.Fail()
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	var TestCases = []TestCase{
		{"aa", true},
		{"aba", true},
		{"eabae", true},
		{"abc", false},
		{"bc", false},
		{"abca", false},
		{"caba", false},
	}

	for _, tc := range TestCases {
		t.Run(fmt.Sprintf("input %q", tc.s), func(t *testing.T) {
			got := isPalindrome(tc.s)
			if got != tc.answer {
				t.Logf("wrong answer in input %q\n", tc.s)
				t.Logf("want %v got %v\n", tc.answer, got)
				t.Fail()
			}
		})
	}
}
