func maxProduct(nums []int) int {
    /*
    动态规划
    这题是求数组中子区间的最大乘积，对于乘法，我们需要注意，负数乘以负数，会变成正数，所以解这题的时候我们需要维护两个变量，当前的最大值，以及最小值，最小值可能为负数，但没准下一步乘以一个负数，当前的最大值就变成最小值，而最小值则变成最大值。 dp[i]代表第i个元素结尾的连续子数组乘积
    1 转移方程
    maxDp[i]=max(maxDp[i-1]*ai,ai,minDp[i-1]*ai)
    minDp[i]=min(minDp[i-1]*ai,ai,maxDp[i-1]*ai)
    Dp[i]=max(Dp[i],maxDp[i])
    2 初始化
    */
    var m=len(nums)
    if m<1 {
        return 0
    }
    if m==1 {
        return nums[0]
    }
    var finaMax=nums[0]
    var maxDp=make([]int,m)
    var minDp=make([]int,m)
    maxDp[0]=nums[0]
    minDp[0]=nums[0]
    for i:=1;i<m;i++ {
        maxDp[i]=max(maxDp[i-1]*nums[i],nums[i],minDp[i-1]*nums[i])
        minDp[i]=min(minDp[i-1]*nums[i],nums[i],maxDp[i-1]*nums[i])
        finaMax=max(finaMax,maxDp[i])
    }
    return finaMax
}
//...接受变长的参数为切片
func max(x ...int)int {
    sort.Ints(x)
    return x[len(x)-1]
}

func min(x ...int)int {
    sort.Ints(x)
    return x[0]
}
