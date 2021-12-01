package main

func candy(ratings []int) int {
	numberOfChildren := len(ratings)
	candy := make([]int, numberOfChildren)
	//Each child must have at least one candy.
	for i := range candy {
		candy[i] = 1
	}

	//from left
	for i := 1; i < numberOfChildren; i++ {
		if ratings[i] > ratings[i-1] && candy[i] <= candy[i-1] {
			candy[i] = candy[i-1] + 1
		}
	}

	//from right
	for i := numberOfChildren - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candy[i] <= candy[i+1] {
			candy[i] = candy[i+1] + 1
		}

	}

	var count int
	for _, c := range candy {
		count += c
	}
	return count
}
