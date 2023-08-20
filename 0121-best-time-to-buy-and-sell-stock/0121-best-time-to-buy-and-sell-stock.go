func maxProfit(prices []int) int {
    buy, profit := prices[0], 0

    for i := 1; i < len(prices); i++ {
        if prices[i] < buy {
            buy = prices[i]
        }

        if prices[i] - buy > profit {
            profit = prices[i] - buy
        }
    }

    return profit
}