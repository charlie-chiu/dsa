package main

import "testing"

type TestCase struct {
	g, s   []int
	answer int
}

func TestFindContentChildren(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{1, 2}, []int{}, 0},
		{[]int{1, 2, 3}, []int{1, 1}, 1},
		{[]int{1, 2}, []int{1, 2, 3}, 2},
		{[]int{3, 3, 3}, []int{1, 2, 3}, 1},
		{[]int{10, 9, 8, 7}, []int{5, 6, 7, 8}, 2},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := findContentChildren(tc.g, tc.s)
			if got != tc.answer {
				t.Logf("wrong answer in input g: %v, s: %v\n", tc.g, tc.s)
				t.Logf("want %d got %d\n", tc.answer, got)
				t.Fail()
			}
		})
	}
}
