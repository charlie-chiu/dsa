package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	s      string
	answer []int
}

func TestPartitionLabels(t *testing.T) {
	var TestCases = []TestCase{
		{"abc", []int{1, 1, 1}},
		{"ababcbacadefegdehijhklij", []int{9, 7, 8}},
		{"eccbbbbdec", []int{10}},
		{"qvmwtmzzse", []int{1, 1, 4, 2, 1, 1}},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := partitionLabels(tc.s)
			if !reflect.DeepEqual(got, tc.answer) {
				t.Logf("wrong answer in input %q", tc.s)
				t.Logf("want %v got %v", tc.answer, got)
				t.Fail()
			}
		})
	}
}
