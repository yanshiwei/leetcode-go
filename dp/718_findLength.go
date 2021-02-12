func findLength(A []int, B []int) int {
    /*
    动态规划
    dp[i][j]下标i-1结尾为A数组，j-1结尾为B数组，最长CS序列长度为dp[i][j]
    转化公式
    dp[i][j]=dp[i-1][j-1]+1 when a[i-1]==b[j-1]
    初始化
    dp[i][0] dp[0][j]无意义，可以初始化0
    */
    var ma=len(A)
    var mb=len(B)
    var dp=make([][]int,ma+1)//i=0无效
    for i:=range dp{
        dp[i]=make([]int,mb+1)//j=0无效
    }
    var maxLen int
    for i:=1;i<ma+1;i++{
        for j:=1;j<mb+1;j++{
            if A[i-1]==B[j-1]{
                dp[i][j]=dp[i-1][j-1]+1
            }
            maxLen=max(maxLen,dp[i][j])
        }
    }
    return maxLen
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
