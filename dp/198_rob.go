func rob(nums []int) int {
    //dp,seen https://leetcode-cn.com/problems/house-robber/solution/dong-tai-gui-hua-jie-ti-si-bu-zou-xiang-jie-cjavap/
    var res int
    if len(nums)<1{
        return res
    }
    var n=len(nums)
    var dp=make([]int,n+1)
    dp[0]=0//0 house
    dp[1]=nums[0]//1 house
    for k:=2;k<=n;k++{
        dp[k]=max(dp[k-1],nums[k-1]+dp[k-2])
    }
    return dp[n]
}

func max(x,y int)int {
    if x<y{
        return y
    }
    return x
}
