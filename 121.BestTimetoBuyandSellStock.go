func maxProfit(prices []int) int {
    n := len(prices)
    if n == 0 {return 0}
    min_prices := make([]int, n)
    max_profit := make([]int, n)
    min_prices[0] = prices[0]
    max_profit[0] = 0
    for i := 1; i < n; i++ {
        min_prices[i] = min(min_prices[i-1], prices[i])
        max_profit[i] = max(max_profit[i-1], prices[i] - min_prices[i-1])
    }
    return max_profit[n-1]
 }

func max(a, b int) int {
    if a < b{
        return b
    }else {
        return a
    }
    
}

func min(a,b int) int {
    if a < b {
        return a
    }else {
        return b
    }
}