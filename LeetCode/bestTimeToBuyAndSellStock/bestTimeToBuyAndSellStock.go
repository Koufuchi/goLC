package bestTimeToBuyAndSellStock

func maxProfit(prices []int) int {
	var max, left int
	for right := 1; right < len(prices); right++ {
		if prices[left] > prices[right] {
			left = right
		} else if prices[right]-prices[left] > max {
			max = prices[right] - prices[left]
		}
	}
	return max
}
