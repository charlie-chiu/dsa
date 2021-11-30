package main

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	var compositeMaker = make([]bool, n+1)
	compositeMaker[0] = true
	compositeMaker[1] = true

	for i := 2; i <= n; i++ {
		if compositeMaker[i] == true {
			continue
		}

		for j := i + i; j <= n; j += i {
			compositeMaker[j] = true
		}
	}

	var count int
	for i := 0; i < n; i++ {
		if compositeMaker[i] == false {
			count++
		}
	}

	return count
}
