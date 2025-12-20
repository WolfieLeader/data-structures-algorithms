package main

// TODO:

func maxProfit(prices []int) int {
	bestProfit := 0

	bestBuy := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < bestBuy {
			bestBuy = prices[i]
			continue
		}
		bestProfit = max(bestProfit, prices[i]-bestBuy)
	}
	return bestProfit
}
