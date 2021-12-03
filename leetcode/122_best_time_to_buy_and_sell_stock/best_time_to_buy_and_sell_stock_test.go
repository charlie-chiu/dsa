package main

import "testing"

type TestCase struct {
	prices []int
	answer int
}

func TestMaxProfit(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{999}, 0},
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 5, 3, 1}, 0},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := maxProfit(tc.prices)
			if got != tc.answer {
				t.Logf("wrong answer in prices %v", tc.prices)
				t.Logf("want %d got %d", tc.answer, got)
				t.Fail()
			}
		})
	}
}
