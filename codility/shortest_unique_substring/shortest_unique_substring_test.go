package main

import "testing"

type TestCase struct {
	input  string
	answer int
}

func TestSolution(t *testing.T) {
	var TestCases = []TestCase{
		{"a", 1},            //a
		{"aa", 2},           //aa
		{"abaaba", 2},       //aa
		{"zyzyzyz", 5},      //zyzyzyz
		{"aabbbabaaa", 3},   //???
		{"aabbbabaaaxx", 2}, //???
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := solution(tc.input)
			if got != tc.answer {
				t.Errorf("wrong answer in input %q, want %d got %d", tc.input, tc.answer, got)
			}
		})
	}
}
