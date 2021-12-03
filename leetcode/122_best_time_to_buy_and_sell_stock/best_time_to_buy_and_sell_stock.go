package main

func maxProfit(prices []int) int {
	totalProfit := 0
	for i := 0; i < len(prices)-1; i++ {
		profit := prices[i+1] - prices[i]
		if profit > 0 {
			totalProfit += profit
		}
	}

	return totalProfit
}
