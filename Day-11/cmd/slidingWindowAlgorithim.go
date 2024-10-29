package main

import (
	"math"
)

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
*/

func maxProfit(prices []int) int {

	// Initialize minPrice to the highest possible value for an unsigned integer,
	// so the first price in the list will replace it. This way, we always keep track of the minimum price seen so far.
	minProfit := math.MaxUint32

	// Initialize maxProfit to 0, since no profit is made initially.
	// This will store the maximum profit we find during the loop.
	maxProfit := 0

	// loop through the prices
	for _, price := range prices {
		// If the current price is higher than minPrice, calculate profit by selling at the current price.
		// Compare this calculated profit with maxProfit and update maxProfit if it's higher.
		if price > minProfit {
			if price-minProfit > maxProfit {
				maxProfit = price - minProfit
			}
		} else {
			// If the current price is less than or equal to minPrice,
			// update minPrice to this price as it's the new lowest buying price.
			minProfit = price
		}
	}

	// After the loop, maxProfit contains the maximum profit found (or 0 if no profit is possible).
	return maxProfit
}
