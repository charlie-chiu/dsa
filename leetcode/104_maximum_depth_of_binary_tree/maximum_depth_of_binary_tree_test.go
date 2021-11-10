package main

import (
	"reflect"
	"testing"
)
import . "dsa/leetcode"

type TestCase struct {
	input  *TreeNode
	answer int
}

func TestMaxDepth(t *testing.T) {
	var TestCases = []TestCase{
		{nil, 0},
		{NewTree(0), 1},
		{NewTree(3, 9, 20, nil, nil, 15, 7), 3},
		{NewTree(1, nil, 2), 2},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := maxDepth(tc.input)
			if !reflect.DeepEqual(got, tc.answer) {
				t.Logf("Wrong answer in input %v\n", tc.input)
				t.Logf("want : %v\n", tc.answer)
				t.Logf(" got : %v\n", got)
				t.Error()
			}
		})
	}
}
