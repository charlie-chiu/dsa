package main

import "testing"

type TestCase struct {
	points [][]int
	answer int
}

func TestShots(t *testing.T) {
	var TestCases = []TestCase{
		{[][]int{{10, 16}}, 1},
		{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, 2},
		{[][]int{{1, 2}, {5, 6}, {3, 4}, {7, 8}}, 4},
		{[][]int{{7, 8}, {1, 8}, {2, 8}, {4, 8}}, 1},
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 2},
		{[][]int{{2, 4}, {1, 4}, {1, 5}, {3, 6}, {7, 10}}, 2},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := findMinArrowShots(tc.points)
			if got != tc.answer {
				t.Logf("wrong answer in input %v", tc.points)
				t.Logf("want %d got %d", tc.answer, got)
				t.Fail()
			}
		})
	}
}
