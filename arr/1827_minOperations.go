func minOperations(nums []int) int {
    if len(nums)<2{
        return 0
    }
    var res int
    for i:=1;i<len(nums);i++{
        if nums[i]<=nums[i-1]{
            res+=nums[i-1]+1-nums[i]
            nums[i]=nums[i-1]+1
        }
    }
    return res
}
