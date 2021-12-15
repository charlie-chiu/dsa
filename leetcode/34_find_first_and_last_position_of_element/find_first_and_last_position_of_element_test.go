package main

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	nums   []int
	target int
	answer []int
}

func TestSearchRange(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{}, 0, []int{-1, -1}},
		{[]int{1}, 1, []int{0, 0}},
		{[]int{5, 7}, 5, []int{0, 0}},
		{[]int{5, 7, 8, 8, 8, 10}, 5, []int{0, 0}},
		{[]int{5, 7, 8, 8, 8, 10}, 7, []int{1, 1}},
		{[]int{5, 7, 8, 8, 8, 10}, 10, []int{5, 5}},
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 10}, 8, []int{3, 8}},
		{[]int{1, 3, 3, 5, 5, 9}, 3, []int{1, 2}},
		{[]int{5, 7, 7, 8, 8, 10}, 3, []int{-1, -1}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{5, 7, 7, 8, 8, 10}, 100, []int{-1, -1}},
	}

	for _, tc := range TestCases {
		t.Run(fmt.Sprintf("%v, %d", tc.nums, tc.target), func(t *testing.T) {
			got := searchRange(tc.nums, tc.target)
			if !reflect.DeepEqual(got, tc.answer) {
				t.Logf("wrong answer in input %v, %v\n", tc.nums, tc.target)
				t.Logf("want %v got %v\n", tc.answer, got)
				t.Fail()
			}
		})
	}
}
