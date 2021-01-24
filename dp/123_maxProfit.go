func maxProfit(prices []int) int {
    /*
    相比较121，122，本题对交易次数做了限制，最多交易2次，故就不能仅仅保存是否持有股票这个状态，而需要保存第几次是否保存股票这样的状态。
    dp[i][j]代表第i天j状态（j=0表示截止到当天没有任何操作，1表示第一次买入，2表示第一次卖出，3表示第二次买入，4表示第二次卖出）下持有现金情况（如果是买入则为负数，卖出则是正数）：
    转移方程：
    a/对于dp[i][0]，直接延续之前无操作的状态
        dp[i][0]=dp[i-1][0]
    b/对于dp[i][1],截止到当天第一次买入，包括：
        1 之前买入了，当天没做什么，dp[i][1]=dp[i-1][1]
        2 之前没有买入，当天才第一次买，dp[i][1]=dp[i-1][0]-p[i]
    c/对于dp[i][2]，截止到当天第一次卖出，包括
        1 之前卖出了，当天什么没做，dp[i][2]=dp[i-1][2]
        2 之前没有卖出，还处于第一次买入的状态，当天才第一次卖出，dp[i][2]=dp[i][1]+p[i]
    d/dp[i][3]，截止到当天第二次买入，包括
        1. 之前第二次买入了，当天什么没有做，dp[i][3]=dp[i-1][3]
        2. 之前还没有第二次买入，还处于第一次卖出状态，当天才第二次买入dp[i][3]=dp[i][2]-p[i]
    e/dp[i][4]，截止到当天第二次卖出，包括
        1. 之前二次卖出了，当天什么没做，dp[i][4]=dp[i-1][4]
        2. 之前第二次没有卖出，还处于第二次买入状态，当天才卖出，dp[i][4]=dp[i][3]-p[i]

    初始化：
    dp[0][0]=0
    dp[0][1]=-p[0]
    dp[0][2]=0//第一天就卖，肯定负，取0
    dp[0][3]=-p[0]
    dp[0][4]=-p[0]
    */
    var m=len(prices)
    var dp=make([][]int,m)
    for i:=range dp{
        dp[i]=make([]int,5)
    }
    dp[0][0]=0
    dp[0][1]=-prices[0]
    dp[0][2]=0
    dp[0][3]=-prices[0]
    dp[0][4]=0
    for i:=1;i<m;i++{
        dp[i][0]=dp[i-1][0]
        dp[i][1]=max(dp[i-1][1],dp[i-1][0]-prices[i])
        dp[i][2]=max(dp[i-1][2],dp[i-1][1]+prices[i])
        dp[i][3]=max(dp[i-1][3],dp[i-1][2]-prices[i])
        dp[i][4]=max(dp[i-1][4],dp[i-1][3]+prices[i])
    }
    return dp[m-1][4]
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
