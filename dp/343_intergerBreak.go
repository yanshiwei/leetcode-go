func integerBreak(n int) int {
    // https://leetcode-cn.com/problems/integer-break/solution/zheng-shu-chai-fen-by-leetcode-solution/
    var dp=make([]int,n+1)//dp[i] 表示将正整数 i 拆分成至少两个正整数的和之后，这些正整数的最大乘积
    //0 不是正整数，1 是最小的正整数，0 和 1 都不能拆分，因此 dp[0]=dp[1]=0
    // init
    dp[0]=0
    dp[1]=1
    for i:=2;i<=n;i++{
        var curMax int
        // 1<=j<i, j是第一个拆分的正整数
        for j:=1;j<i;j++{
            first:=j*(i-j)//i-j不拆分
            second:=j*dp[i-j]//i-j继续拆分
            final:=max(first,second)
            if curMax<final{
                curMax=final
            }
        }
        // loop all j and get max
        dp[i]=curMax
    }
    return dp[n]
}

func max(x,y int)int {
    if x<y{
        return y
    }
    return x
}
