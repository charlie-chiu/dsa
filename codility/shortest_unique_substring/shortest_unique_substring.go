package main

func solution(s string) int {
	//var shortestLength int
	var subCounter = make(map[string]int)

	for substringLen := 1; substringLen <= len(s); substringLen++ {
		for i := 0; i <= len(s)-substringLen; i++ {
			subStr := s[i : i+substringLen]
			subCounter[subStr]++
			//log.Printf("start at %d, length %d, subStr %s",i, substringLen, subStr)
		}

		for _, i := range subCounter {
			if i == 1 {
				return substringLen
			}
		}

		subCounter = map[string]int{}
	}

	return 0
}
