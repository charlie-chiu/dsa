package main

func lengthOfLongestSubstring(s string) int {
	var start, end, globalMax int
	var charCount = make(map[byte]int)

	for end < len(s) {
		charCount[s[end]]++
		for start <= end && hasRepeated(charCount) {
			charCount[s[start]] -= 1
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

func hasRepeated(counter map[byte]int) bool {
	for _, count := range counter {
		if count > 1 {
			return true
		}
	}

	return false
}
