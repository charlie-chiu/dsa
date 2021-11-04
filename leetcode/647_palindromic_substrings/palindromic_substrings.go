package main

func countSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		count += countHelper(s, i, i)
		count += countHelper(s, i, i+1)
	}

	return count
}

func countHelper(s string, l, r int) int {
	var count = 0

	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
		count++
	}

	return count
}
