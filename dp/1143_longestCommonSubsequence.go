func longestCommonSubsequence(text1 string, text2 string) int {
    /*
    动态规划，类似718 https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/
    dp[i][j]下标i-1结尾为A数组，j-1结尾为B数组，最长CS序列长度为dp[i][j]
    转化公式
    dp[i][j]=dp[i-1][j-1]+1 when a[i-1]==b[j-1]
    dp[i][j]=max(dp[i-1][j],dp[i-1][j])
    初始化
    dp[i][0] dp[0][j]无意义，可以初始化0
    */
    var ma=len(text1)
    var mb=len(text2)
    var dp=make([][]int,ma+1)//i=0无效
    for i:=range dp{
        dp[i]=make([]int,mb+1)//j=0无效
    }
    var maxLen int
    for i:=1;i<ma+1;i++{
        for j:=1;j<mb+1;j++{
            if text1[i-1]==text2[j-1]{
                dp[i][j]=dp[i-1][j-1]+1
            }else{
                dp[i][j]=max(dp[i-1][j],dp[i][j-1])
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
