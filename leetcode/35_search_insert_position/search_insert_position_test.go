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

func TestSearchInsert(t *testing.T) {
	var TestCases = []TestCase{
		// found
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6, 10, 20, 50, 999}, 3, 1},
		{[]int{1, 3, 5, 6, 10, 20, 50, 999}, 50, 6},

		//insert
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4}, // insert at last one
	}

	for _, tc := range TestCases {
		t.Run(fmt.Sprintf("nums %v target %d", tc.nums, tc.target), func(t *testing.T) {
			got := searchInsert(tc.nums, tc.target)
			if got != tc.answer {
				t.Logf("wrong answer in input %v, %d\n", tc.nums, tc.target)
				t.Logf("want %d got %d\n", tc.answer, got)
				t.Fail()
			}
		})
	}
}
