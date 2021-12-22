package main

import (
	"fmt"
	"testing"
)

type TestCase struct {
	nums   []int
	answer int
}

func TestSearch(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{1}, 1},
		{[]int{5}, 5},
		{[]int{5, 9, 9}, 5},
		{[]int{1, 1, 2}, 2},
		{[]int{1, 1, 2, 3, 3, 4, 4, 8, 8}, 2},
		{[]int{3, 3, 7, 7, 10, 11, 11}, 10},
	}

	for _, tc := range TestCases {
		t.Run(fmt.Sprintf("finding in %v", tc.nums), func(t *testing.T) {
			got := singleNonDuplicate(tc.nums)
			if got != tc.answer {
				t.Logf("wrong answer in input %v\n", tc.nums)
				t.Logf("want %d got %d\n", tc.answer, got)
				t.Fail()
			}
		})
	}
}
