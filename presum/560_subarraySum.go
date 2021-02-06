func subarraySum(nums []int, k int) int {
    /*
    前缀和数组保存前 n 位的和，presum[1]保存的就是 nums 数组中前 1 位的和，也就是 presum[1] = nums[0], presum[2] = nums[0] + nums[1] = presum[1] + nums[1]. 依次类推
    */
    var res int
    if len(nums)<1 {
        return res
    }
    var pre=make([]int,len(nums)+1)
    for i:=1;i<len(nums)+1;i++ {
        pre[i]=pre[i-1]+nums[i-1]
    }
    for i:=1;i<len(nums)+1;i++ {
        for j:=i;j<len(nums)+1;j++ {
            if (pre[j]-pre[i-1])==k {
                res++
            }
        }
    }
    return res
}
