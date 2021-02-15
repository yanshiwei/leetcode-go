func minCostClimbingStairs(cost []int) int {
    /*
    cost 的长度为 n，则 n 个阶梯分别对应下标 00 到 n-1，楼层顶部对应下标 n，问题等价于计算达到下标 n的最小花费。可以通过动态规划求解
    创建长度n+1的数组dp，dp[i]表示到达下标i所需的最小消费
    转移方程：
    dp[i]=min(dp[i-1]+cost[i-1],dp[i-2]+cost[i-2])
    初始化：
    可以选择下标 0 或 1作为初始阶梯故dp[0]=0 dp[1]=0
    */
    var m=len(cost)
    var dp=make([]int,m+1)
    for i:=2;i<m+1;i++{
        dp[i]=min(dp[i-1]+cost[i-1],dp[i-2]+cost[i-2])
    }
    return dp[m]
}
func min(x,y int)int {
    if x<y{
        return x
    }
    return y
}
