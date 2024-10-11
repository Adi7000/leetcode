package leetcode

func maxProfit(prices []int, fee int) int {
	lastDay := len(prices)
	if lastDay < 2 {
		return 0
	}
	buyingPrice := prices[0] + 1
	var sellingPrice *int
	profit := 0

	for currentDay := 0; currentDay < lastDay; currentDay++ {
		currentDayPrice := prices[currentDay]

		if sellingPrice == nil && currentDayPrice < buyingPrice {
			//potential buying price
			buyingPrice = currentDayPrice
		} else if sellingPrice != nil && currentDayPrice < *sellingPrice-fee {
			//buy and sell
			profit += *sellingPrice - buyingPrice - fee
			sellingPrice = nil
			buyingPrice = currentDayPrice
		} else if (sellingPrice == nil && currentDayPrice > buyingPrice+fee) ||
			(sellingPrice != nil && currentDayPrice > *sellingPrice) {
			//potential selling price
			sellingPrice = &currentDayPrice
		}
	}
	if sellingPrice != nil {
		return max(profit, profit+*sellingPrice-buyingPrice-fee)
	}
	return profit

}
