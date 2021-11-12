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

func TestDiameterOfBinaryTree(t *testing.T) {
	var TestCases = []TestCase{
		{NewTree(0), 0},
		{NewTree(1, 2, 3, 4, 5), 3},
		{NewTree(1, 2), 1},
		{NewTree(9, 6, nil, 0, 6, nil, 1, 4), 4},
		{NewTree(4, -7, -3, nil, nil, -9, -3, 9, -7, -4, nil, 6, nil, -6, -6, nil, nil, 0, 6, 5, nil, 9, nil, nil, -1, -4, nil, nil, nil, -2), 8},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := diameterOfBinaryTree(tc.input)
			if !reflect.DeepEqual(got, tc.answer) {
				t.Logf("Wrong answer in input %v\n", tc.input)
				t.Logf("want : %v\n", tc.answer)
				t.Logf(" got : %v\n", got)
				t.Error()
			}
		})
	}
}
