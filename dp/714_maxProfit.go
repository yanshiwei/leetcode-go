func maxProfit(prices []int, fee int) int {
    /*
    动态规划
    dp[i][j]代表第i天j状态（j=0表示当天不持有股票，=1则表示持有）下持有现金情况（如果是买入则为负数，卖出则是正数）：
    转移方程
    1 第i天不持有股票，则有如下两种情况：
        1.1 昨天不持有，今天不做什么
        1.2 昨天持有，今天卖出
        dp[i][0]=max(dp[i-1][0],dp[i-1][1]+p[i]-fee)
    2 第i天持有股票，则有如下两种情况：
        2.1 昨天持有，今天不做什么
        2.2 昨天没有，今天买入
        dp[i][1]=max(dp[i-1][1],dp[i-1][0]-p[i])
    初始状态
    dp[0][0]=0//不持股显然为 0
    dp[1][0]=-p[0]//第一天就买入
    */
    var res int
    if len(prices)<2 {
        return res
    }
    var m=len(prices)
    var dp=make([][]int,m)
    for i:=range dp{
        dp[i]=make([]int,2)
        //两种状态
    }
    //初始化
    dp[0][0]=0
    dp[0][1]=-prices[0]
    for i:=1;i<m;i++{
        dp[i][0]=max(dp[i-1][0],dp[i-1][1]+prices[i]-fee)
        dp[i][1]=max(dp[i-1][1],dp[i-1][0]-prices[i])
    }
    return dp[m-1][0]
}

func max(x, y int)int {
    if x<y {
        return y
    }
    return x
}
