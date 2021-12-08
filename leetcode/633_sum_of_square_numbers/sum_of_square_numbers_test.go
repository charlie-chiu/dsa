package main

import (
	"fmt"
	"testing"
)

type TestCase struct {
	c      int
	answer bool
}

func TestJudgeSquareSum(t *testing.T) {
	var TestCases = []TestCase{
		{0, true},
		{1, true},
		{2, true},
		{3, false},
		{4, true},
		{5, true},
		{6, false},
		{7, false},
		{8, true},
		{2147483641, false},
	}

	for _, tc := range TestCases {
		t.Run(fmt.Sprintf("input:%d", tc.c), func(t *testing.T) {
			got := judgeSquareSum(tc.c)
			if got != tc.answer {
				t.Logf("wrong answer in input %v\n", tc.c)
				t.Logf("want %v got %v\n", tc.answer, got)
				t.Fail()
			}
		})
	}
}
