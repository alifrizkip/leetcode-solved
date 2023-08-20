func maxProfit(prices []int) int {
    buy, sell := prices[0], 0
    for _, p := range prices {
        if p < buy {
            buy = p
        }

        if p - buy > sell {
            sell = p - buy
        }
    }

    return sell
}