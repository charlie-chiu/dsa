package main

//first try, 8ms, 5.6MB on leetcode
func countBinarySubstrings(s string) int {
	if len(s) < 2 {
		return 0
	}

	count := 0

	for i := 0; i < len(s); i++ {
		l := i
		r := i + 1
		for l >= 0 && r < len(s) && s[l] != s[r] && !(r-l > 1 && s[l] != s[l+1]) {
			l--
			r++
			count++
		}
	}

	return count
}
