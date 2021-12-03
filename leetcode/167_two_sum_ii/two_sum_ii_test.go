package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	numbers []int
	target  int
	answer  []int
}

func TestTwoSum(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{-4, 4, 9, 10, 15}, 14, []int{2, 4}},
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := twoSum(tc.numbers, tc.target)
			if !reflect.DeepEqual(got, tc.answer) {
				t.Logf("wrong answer in numbers: %v target: %d\n", tc.numbers, tc.target)
				t.Logf("want %v got %v\n", tc.answer, got)
				t.Fail()
			}
		})
	}
}
