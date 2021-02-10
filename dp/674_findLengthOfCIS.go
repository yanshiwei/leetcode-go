func findLengthOfLCIS(nums []int) int {
    var res int
    //以下标i为结尾的数组的连续递增的子序列长度为dp[i]
    //公式 dp[i + 1] = dp[i] + 1 if num[i+1]>num[i]
    //初始化 ，标i为结尾的数组的连续递增的子序列长度最少也应该是 故dp[i]=1 i是任意值
    if len(nums)<1{
        return res
    }
    if len(nums)==1{
        return 1
    }
    var dp=make([]int,len(nums))
    for i:=range dp{
        dp[i]=1//连续递增的子序列长度最少也应该是1
    }
    for i:=0;i<len(nums)-1;i++{
        if nums[i+1]>nums[i]{
            dp[i+1]=dp[i]+1
        }
        res=max(dp[i+1],res)
    }
    return res
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
