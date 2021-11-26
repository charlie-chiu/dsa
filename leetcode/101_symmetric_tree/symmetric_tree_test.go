package main

import (
	"reflect"
	"testing"

	. "dsa/leetcode"
)

type TestCase struct {
	input  *TreeNode
	answer bool
}

func TestIsBalanced(t *testing.T) {
	var TestCases = []TestCase{
		{nil, true},
		{NewTree(9), true},
		{NewTree(1, 2, 2, 3, 4, 4, 3), true},
		{NewTree(1, 2, 2, nil, 3, nil, 3), false},

		// doesn't know why leetcode don't have test case like this
		//{NewTree(1,2,2,3,3,3,3,nil,nil,4,nil,nil,4), false},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := isSymmetric(tc.input)
			if !reflect.DeepEqual(got, tc.answer) {
				t.Logf("Wrong answer in input %v\n", tc.input)
				t.Logf("want : %v\n", tc.answer)
				t.Logf(" got : %v\n", got)
				t.Error()
			}
		})
	}
}
