package main

import "testing"

type TestCase struct {
	input  int
	answer string
}

func TestConvertToBase7(t *testing.T) {
	var TestCases = []TestCase{
		{0, "0"},
		{3, "3"},
		{-7, "-10"},
		{7, "10"},
		{100, "202"},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := convertToBase7(tc.input)
			if got != tc.answer {
				t.Logf("wrong answer in input %d", tc.input)
				t.Logf("want %q got %q", tc.answer, got)
				t.Fail()
			}
		})
	}
}
