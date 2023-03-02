package day7

import "math"

//^股票最大利润问题
func maxProfit(prices []int) int {
	cost, profit := math.MaxInt, 0 //?买入价格，利润
	for i := 0; i < len(prices); i++ {
		if prices[i] < cost {
			cost = prices[i]
		}

		p := prices[i] - cost
		if p > profit {
			profit = p
		}
	}
	return profit
}
