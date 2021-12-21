package main

import (
	"fmt"
	"testing"
)

type TestCase struct {
	nums   []int
	target int
	answer bool
}

func TestSearch(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{0}, 1, false},
		{[]int{0}, 0, true},
		{[]int{1, 3, 5, 6, 7, 8}, 1, true},
		{[]int{1, 3, 5, 6, 7, 8}, 4, false},

		//with duplicated
		{[]int{1, 3, 3, 3, 5, 6, 6, 6, 7, 8}, 4, false},
		{[]int{1, 3, 3, 3, 5, 6, 6, 6, 7, 8}, 9, false},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 3, 3, 5, 6, 6, 6, 7, 8}, 1, true},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 3, 3, 5, 6, 6, 6, 7, 8}, 3, true},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 3, 3, 5, 6, 6, 6, 7, 8}, 7, true},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 3, 3, 5, 6, 6, 6, 7, 8}, 8, true},

		// leetcode test cases
		{[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 3, false},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}, 2, true},
	}

	for _, tc := range TestCases {
		t.Run(fmt.Sprintf("finding %d", tc.target), func(t *testing.T) {
			got := search(tc.nums, tc.target)
			if got != tc.answer {
				t.Logf("wrong answer in input %v %v\n", tc.nums, tc.target)
				t.Logf("want %v got %v\n", tc.answer, got)
				t.Fail()
			}
		})
	}
}
