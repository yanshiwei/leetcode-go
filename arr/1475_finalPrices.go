func finalPrices(prices []int) []int {
    for i:=0;i<len(prices);i++{
        var j int
        for j=i+1;j<len(prices);j++{
            if prices[j]<=prices[i]{
                break
            }
        }
        if j>0&&j<len(prices) {
            prices[i]=prices[i]-prices[j]
        }
    }
    return prices
}
