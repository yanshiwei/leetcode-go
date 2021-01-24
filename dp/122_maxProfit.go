func maxProfit(prices []int) int {
/*
与121类似，但是121是只交易一次，该题允许交易多次取和，只不过交易顺序有先后即必须先买入卖出第一个才能买入卖出下一个。
动态规划
    dp[i][j]代表第i天j状态（j=0表示当天不持有股票，=1则表示持有）下持有现金情况（如果是买入则为负数，卖出则是正数）：
    转移方程：
    1第i天不持有股票，则存在两种情况：
    1.1 昨天不持有，当日不做什么
    1.2 昨天持有，当日卖出
    dp[i][0]=max(dp[i-1][0],dp[i-1][1]+p[i])
    2第i天持有股票，也存在两种情况：
    2.1 昨天持有，今天不做什么
    2.2 昨天没有，今天花钱买入
    121题是dp[i][1]=max(dp[i-1][1],-p[i])
    考虑到多次交易，故还需要考虑历史现金dp[i-1][0],之所以是dp[i-1][0]而不是dp[i-1][1]，是题目要求先买入卖出第一个才能买入卖出下一个故dp[i][1]=max(dp[i-1][1],dp[i-1][0]-p[i])
    初始化
    dp[0][0]=0//不持股显然为 0
    dp[1][0]=-p[0]//第一天就买入
*/
    if len(prices)<2 {
        return 0
    }
    var m=len(prices)
    var dp=make([][]int,m)
    for i:=range dp{
        dp[i]=make([]int,2)
    }
    dp[0][0]=0
    dp[0][1]=-prices[0]
    for i:=1;i<m;i++ {
        dp[i][0]=max(dp[i-1][0],dp[i-1][1]+prices[i])
        dp[i][1]=max(dp[i-1][1],dp[i-1][0]-prices[i])
    }
    return dp[m-1][0]
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
