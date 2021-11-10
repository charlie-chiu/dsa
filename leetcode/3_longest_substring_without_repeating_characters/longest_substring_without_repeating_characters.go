package main

func lengthOfLongestSubstring(s string) int {
	var start, end, globalMax int
	// length 128 to handle whole ascii table
	var countSlice = make([]int, 128)

	for end < len(s) {
		countSlice[s[end]]++
		for start <= end && countSlice[s[end]] > 1 {
			countSlice[s[start]] -= 1
			start++
		}

		globalMax = max(globalMax, end-start+1)
		end++
	}

	return globalMax
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
