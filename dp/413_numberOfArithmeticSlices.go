func numberOfArithmeticSlices(nums []int) int {
    var res int
    if len(nums)<3{
        return res
    }
    var dp=make([]int,len(nums))//dp[i]表示以nums[i]为结尾的等差数列数量
    for i:=2;i<len(nums);i++{
        if nums[i]-nums[i-1]==nums[i-1]-nums[i-2]{
            dp[i]=dp[i-1]+1
            res+=dp[i]
        }else{
            dp[i]=0
        }
    }
    return res
}
