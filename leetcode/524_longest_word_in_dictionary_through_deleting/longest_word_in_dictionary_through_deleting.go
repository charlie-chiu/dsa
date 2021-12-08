package main

import (
	"sort"
)

func findLongestWord(s string, dictionary []string) string {
	// sort by word length ,we can start from last, to find the longest match
	sort.SliceStable(dictionary, func(i, j int) bool {
		return len(dictionary[i]) > len(dictionary[j])
	})

	for i := 0; i < len(dictionary); i++ {
		dict := dictionary[i]
		if dict == s || isContain(dict, s) {
			return dict
		}
	}

	return ""
}

func isContain(dict, s string) bool {
	var iS, iDict int
	for iDict < len(dict) {
		if iS >= len(s) {
			return false
		} else if dict[iDict] == s[iS] {
			iS++
			iDict++
		} else {
			iS++
		}
	}

	return true
}
