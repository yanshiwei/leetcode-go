func maxUncrossedLines(A []int, B []int) int {
    /*
    动态规划，与最长公共子序列长1143 一样。
    https://leetcode-cn.com/problems/longest-common-subsequence/solution/
    */
    var ma=len(A)
    var mb=len(B)
    var dp=make([][]int,ma+1)
    for i:=range dp{
        dp[i]=make([]int,mb+1)
    }
    var maxLen int
    for i:=1;i<ma+1;i++{
        for j:=1;j<mb+1;j++{
            if A[i-1]==B[j-1]{
                dp[i][j]=dp[i-1][j-1]+1
            }else{
                dp[i][j]=max(dp[i-1][j],dp[i][j-1])
            }
            maxLen=max(maxLen,dp[i][j])
        }
    }
    return maxLen
}
func max(x,y int)int{
    if x<y {
        return y
    }
    return x
}
