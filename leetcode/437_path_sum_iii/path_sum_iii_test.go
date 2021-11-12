package main

import (
	"reflect"
	"testing"
)
import . "dsa/leetcode"

type TestCase struct {
	input             *TreeNode
	targetSum, answer int
}

func TestDiameterOfBinaryTree(t *testing.T) {
	var TestCases = []TestCase{
		{NewTree(1, 2, 3), 3, 2},
		{NewTree(10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1), 8, 3},
		{NewTree(5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1), 22, 3},
		{NewTree(1, -2, -3, 1, 3, -2, nil, -1), -1, 4},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := pathSum(tc.input, tc.targetSum)
			if !reflect.DeepEqual(got, tc.answer) {
				t.Logf("Wrong answer in input %v, target: %d\n", tc.input, tc.targetSum)
				t.Logf("want : %v\n", tc.answer)
				t.Logf(" got : %v\n", got)
				t.Error()
			}
		})
	}
}
