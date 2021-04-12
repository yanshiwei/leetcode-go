func minStartValue(nums []int) int {
    var minV =nums[0]
    for i:=range nums{
        if i==0{
            continue
        }
        nums[i]=nums[i]+nums[i-1]
        if minV>nums[i]{
            minV=nums[i]
        }
    }
    if minV<0 {
        return 1-minV
    }
    return 1
}
