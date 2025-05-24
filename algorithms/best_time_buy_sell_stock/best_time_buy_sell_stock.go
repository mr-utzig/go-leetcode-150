package best_time_buy_sell_stock

// Example 1:
//
//	Input: prices = [7,1,5,3,6,4]
//	Output: 5

// Example 2:
//
//	Input: prices = [7,6,4,3,1]
//	Output: 0

func Algorithm(prices []int) int {
	cheapest := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < cheapest {
			cheapest = prices[i]
		} else {
			sell := prices[i] - cheapest

			if sell > profit {
				profit = sell
			}
		}
	}

	return profit
}
