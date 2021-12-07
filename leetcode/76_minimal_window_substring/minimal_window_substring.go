package main

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	flag := make([]bool, 128)
	chars := make([]int, 128)

	for i := 0; i < len(t); i++ {
		flag[t[i]] = true
		chars[t[i]]++
	}

	var cnt, l, minL int
	var minSize = len(s) + 1

	for r := 0; r < len(s); r++ {
		if flag[s[r]] {
			chars[s[r]]--
			if chars[s[r]] >= 0 {
				cnt++
			}

			for cnt == len(t) {
				if r-l+1 < minSize {
					minL = l
					minSize = r - l + 1
				}
				if flag[s[l]] == true {
					chars[s[l]]++
					if chars[s[l]] > 0 {
						cnt--
					}
				}
				l++
			}
		}
	}

	if minSize > len(s) {
		return ""
	} else {
		return s[minL : minL+minSize]
	}
}
