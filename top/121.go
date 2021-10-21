package main

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	min, maxPro := prices[0], 0

	for i := 1; i < len(prices); i++ {
		if prices[i]-min > maxPro {
			maxPro = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return maxPro
}
