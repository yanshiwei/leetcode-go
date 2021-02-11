func numSubarrayProductLessThanK(nums []int, k int) int {
    var res int
    if k<=1{
        return res
    }
    var left=0
    var pruduct=1
    for right:=range nums{
        pruduct*=nums[right]
        for pruduct>=k{
            pruduct=pruduct/nums[left]
            left++
        }
        res+=right-left+1
    }
    return res
}
