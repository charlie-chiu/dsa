package main

func longestPalindrome(s string) int {
	var counter = make(map[byte]int)
	var result int

	// count char
	for i := 0; i < len(s); i++ {
		counter[s[i]]++
	}

	// add paired chars
	for char := range counter {
		for counter[char] >= 2 {
			result += 2
			counter[char] -= 2
		}
	}

	// add paired chars
	for _, count := range counter {
		if count > 0 {
			result++
			break
		}
	}

	return result
}
