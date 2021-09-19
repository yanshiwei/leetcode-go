func waysToStep(n int) int {
    //dp
    if n==1{
        return 1
    }
    if n==2{
        return 2
    }
    var dp=make([]int,n+1)
    dp[0]=1
    dp[1]=1
    dp[2]=2
    dp[3]=4
    for i:=4;i<=n;i++{
        dp[i]=(dp[i-1]%1000000007+dp[i-2]%1000000007+dp[i-3]%1000000007)%1000000007
    }
    return dp[n]
}
