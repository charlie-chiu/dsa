package main

func strStr(haystack string, needle string) int {
	if needle == "" || haystack == needle {
		return 0
	}

	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] == needle[j] {
				if j == len(needle)-1 {
					return i
				}
			} else {
				break
			}
		}
	}

	return -1
}
