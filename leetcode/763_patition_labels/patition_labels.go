package main

func partitionLabels(s string) []int {
	// using slice instead map[byte]int to improve performance
	var lastPosition = make([]int, 123)
	for i := 0; i < len(s); i++ {
		lastPosition[s[i]] = max(lastPosition[s[i]], i)
	}

	var positions []int
	var start, curr int
	var end = lastPosition[s[start]]

	for curr < len(s) {
		end = max(end, lastPosition[s[curr]])

		if curr == end {
			positions = append(positions, end-start+1)
			start = end + 1
			end = start
			curr = start
		} else {
			curr++
		}
	}

	return positions
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
