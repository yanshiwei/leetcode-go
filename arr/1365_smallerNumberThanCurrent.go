func smallerNumbersThanCurrent(nums []int) []int {
    var res []int
    if len(nums)<2{
        return res
    }
    for i:=0;i<len(nums);i++{
        var curSum int
        for j:=0;j<len(nums);j++{
            if i==j{
                continue
            }
            if nums[i]>nums[j]{
                curSum++
            }
        }
        res=append(res,curSum)
    }
    return res
}
