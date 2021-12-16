package main

import (
	"fmt"
	"testing"
)

type TestCase struct {
	nums   []int
	target int
	answer int
}

func TestSearch(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{1}, 1, 0},
		{[]int{1}, 0, -1},
		{[]int{1, 3}, 2, -1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{12, 13, 14, 15, 16, 17, 18, 19, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 14, 2},
		{[]int{12, 13, 14, 15, 16, 17, 18, 19, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 7, 15},
	}

	for _, tc := range TestCases {
		t.Run(fmt.Sprintf("nums %v target %d", tc.nums, tc.target), func(t *testing.T) {
			got := search(tc.nums, tc.target)
			if got != tc.answer {
				t.Logf("wrong answer in input %v, %d\n", tc.nums, tc.target)
				t.Logf("want %d got %d\n", tc.answer, got)
				t.Fail()
			}
		})
	}
}
