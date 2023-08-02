func maxProfit(prices []int) int {
    buyPrice, sellPrice := prices[0], 0

    for _, p := range prices {
        if p < buyPrice {
            buyPrice = p
        }

        currSellPrice := p - buyPrice
        if currSellPrice > sellPrice {
            sellPrice = currSellPrice
        }
    }

    return sellPrice
}