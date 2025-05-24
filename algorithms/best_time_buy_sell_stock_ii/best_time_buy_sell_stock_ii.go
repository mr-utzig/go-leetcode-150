package best_time_buy_sell_stock_ii

// Example 1:
// Input: prices = [7,1,5,3,6,4]
// Output: 7

// Example 2:
// Input: prices = [1,2,3,4,5]
// Output: 4

// Example 3:
// Input: prices = [7,6,4,3,1]
// Output: 0

func Algorithm(prices []int) int {
	profit := 0
	l := len(prices)
	for i := 0; i < l; i++ {
		if i+1 < l {
			if prices[i] < prices[i+1] {
				profit += prices[i+1] - prices[i]
			}
		}
	}

	return profit
}
