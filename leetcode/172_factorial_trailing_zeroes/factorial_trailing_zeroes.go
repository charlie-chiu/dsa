package main

func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	}

	return n/5 + trailingZeroes(n/5)
}

//func trailingZeroes(n int) int {
//	var result int
//	for i := 5; i <= n; i *= 5 {
//		result += n / i
//	}
//	return result
//}
