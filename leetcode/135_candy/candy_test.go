package main

import "testing"

type TestCase struct {
	rating []int
	answer int
}

func TestCandy(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{1, 1, 1}, 3},
		{[]int{4, 4, 4}, 3},

		//    2 1 2
		{[]int{1, 0, 2}, 5},
		{[]int{1, 2, 2}, 4},
		{[]int{1, 5, 7, 1}, 7},

		//           1,2,3,2,1
		{[]int{1, 5, 7, 7, 1}, 9},

		//    1 2 3 1 3 2 1
		{[]int{1, 2, 87, 87, 87, 2, 1}, 13},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := candy(tc.rating)
			if got != tc.answer {
				t.Logf("wrong answer in input %v\n", tc.rating)
				t.Logf("want %d got %d\n", tc.answer, got)
				t.Fail()
			}

		})
	}
}
