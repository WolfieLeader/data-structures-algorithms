package main

// HACK: Core Pattern for Greedy

// Greedy
//
// Pattern
// Scan prices once while making locally optimal decisions.
//
// Key idea
// Always keep the lowest price seen so far as the best buy.
// At each day, check the profit if selling today.
// Keep the maximum profit.
//
// Invariant
// buy is the minimum price among all previous days.
// profit is the maximum achievable profit up to today.
//
// Why it works
// Any optimal sell must happen after the lowest possible buy.
// Updating the buy price greedily guarantees the best future profit.
// No past decision needs to be revisited.
//
// Time and space
// Single pass.
// Constant extra space.
//
// Walkthrough: prices = [7,1,5,3,6,4]
// start buy=7 profit=0
// price=1 -> buy=1
// price=5 -> profit=4
// price=3 -> buy=1
// price=6 -> profit=5
// price=4 -> buy=1
// result = 5

func maxProfit(prices []int) int {
	buy, profit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		profit = max(profit, prices[i]-buy)
		buy = min(buy, prices[i])
	}
	return profit
}
