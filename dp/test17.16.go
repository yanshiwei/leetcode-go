func massage(nums []int) int {
    //dp[i] =max(dp[i-2]+nums[i],dp[i-1])
    var n=len(nums)
    if n<1{
        return 0
    }
    if n==1{
        return nums[0]
    }
    var dp=make([]int,n)
    dp[0]=nums[0]
    dp[1]=max(nums[0],nums[1])
    for i:=2;i<n;i++{
        dp[i]=max(dp[i-1],dp[i-2]+nums[i])
    }
    return dp[n-1]
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
