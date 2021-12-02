package main

import "testing"

type TestCase struct {
	flowerbed []int
	n         int
	answer    bool
}

func TestCanPlaceFlowers(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{1}, 0, true},
		{[]int{1}, 1, false},
		{[]int{1}, 9, false},
		{[]int{0}, 1, true},
		{[]int{0}, 4, false},
		{[]int{1, 0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1}, 2, false},
		{[]int{1, 0, 0, 0, 1}, 2, false},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := canPlaceFlowers(tc.flowerbed, tc.n)
			if got != tc.answer {
				t.Logf("wrong answer in input %v, %d", tc.flowerbed, tc.n)
				t.Logf("want %v got %v", tc.answer, got)
				t.Fail()
			}
		})
	}
}
