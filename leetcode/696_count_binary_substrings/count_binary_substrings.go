package main

//first try, 8ms, 5.6MB on leetcode
//func countBinarySubstrings(s string) int {
//	if len(s) < 2 {
//		return 0
//	}
//
//	count := 0
//
//	for i := 0; i < len(s); i++ {
//		l := i
//		r := i + 1
//		for l >= 0 && r < len(s) && s[l] != s[r] && !(r-l > 1 && s[l] != s[l+1]) {
//			l--
//			r++
//			count++
//		}
//	}
//
//	return count
//}

// from leetcode 101, 4ms, 5.6MB
func countBinarySubstrings(s string) int {
	if len(s) < 2 {
		return 0
	}
	count := 0

	var prevLength, currLength int
	currChar := s[0]
	for i := 0; i < len(s); i++ {
		if s[i] != currChar {
			prevLength = currLength
			currLength = 0
			currChar = s[i]
		}

		currLength++

		if prevLength >= currLength {
			count++
		}
	}

	return count
}
