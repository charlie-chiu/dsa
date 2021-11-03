package main

import (
	"reflect"
)

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	return reflect.DeepEqual(transform(t), transform(s))
}

func transform(s string) []int {
	var firstOccurrences = make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := firstOccurrences[s[i]]; !ok {
			firstOccurrences[s[i]] = i
		}
	}
	result := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		result[i] = firstOccurrences[s[i]]
	}

	return result
}
