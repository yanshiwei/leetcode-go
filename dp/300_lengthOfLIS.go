func lengthOfLIS(nums []int) int {
    var n=len(nums)
    var res int
    if n<1{
        return res
    }
    var dp=make([]int,n+1)//仅使用[0,i] (这里是区间的意思，表示前i+1 个)，且以第 i 个结尾时的最大高度
    for i:=0;i<n;i++{
        var maxh int
        for j:=0;j<i;j++{
            if nums[j]<nums[i]{
                if maxh<dp[j]{
                    maxh=dp[j]
                }
            }
        }
        dp[i]=maxh+1
        if res<dp[i]{
            res=dp[i]
        }
    }
    return res
}
