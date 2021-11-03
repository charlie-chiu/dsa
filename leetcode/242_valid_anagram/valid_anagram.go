package main

func isAnagram(s string, t string) bool {
	var counter = make(map[byte]int)

	for i := 0; i < len(s); i++ {
		counter[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		counter[t[i]]--
	}

	for _, diff := range counter {
		if diff != 0 {
			return false
		}
	}

	return true
}
