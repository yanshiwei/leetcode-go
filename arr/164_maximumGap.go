func maximumGap(nums []int) int {
    var res int
    if len(nums)<2{
        return res
    }
    sort.Ints(nums)
    for i:=1;i<len(nums);i++{
        delta:=nums[i]-nums[i-1]
        if delta>res{
            res=delta
        }
    }
    return res
}
