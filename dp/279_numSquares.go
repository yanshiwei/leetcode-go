func numSquares(n int) int {
    //f[i]表示最少需要多少个数的平方来表示整数 i.这些数必然落在区间 [1,sqrt{n}],当枚举到假设当前枚举到 j,//dp. f[i]= min f(i-j^2)+1
    var dp=make([]int,n+1)//dp[0]
    for i:=1;i<=n;i++{
        var minV=INTMAX
        for j:=1;j*j<=i;j++{
            minV=min(minV,dp[i-j*j])
        }
        dp[i]=minV+1
    }
    return dp[n]
}
const INTMAX=int(^uint(0)>>1)
func min(x,y int)int {
    if x<y {
        return x
    }
    return y
}
