package main

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	// not cookie to assign!
	if len(s) == 0 {
		return 0
	}

	sort.Ints(g)
	sort.Ints(s)

	var count, gPointer, sPointer int

	for gPointer < len(g) && sPointer < len(s) {
		if s[sPointer] >= g[gPointer] {
			count++
			sPointer++
			gPointer++
		} else {
			sPointer++
		}
	}

	return count
}
