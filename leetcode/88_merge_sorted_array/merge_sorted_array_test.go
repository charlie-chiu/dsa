package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	nums1  []int
	m      int
	nums2  []int
	n      int
	answer []int
}

func TestMerge(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
		{[]int{0}, 0, []int{2}, 1, []int{2}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
		{[]int{1, 0}, 1, []int{2}, 1, []int{1, 2}},
		{[]int{2, 0}, 1, []int{1}, 1, []int{1, 2}},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			merge(tc.nums1, tc.m, tc.nums2, tc.n)
			if !reflect.DeepEqual(tc.nums1, tc.answer) {
				t.Logf("wrong answer in input %v, %v\n", tc.nums1, tc.nums2)
				t.Logf("want %v got %v\n", tc.answer, tc.nums1)
				t.Fail()
			}
		})
	}
}
