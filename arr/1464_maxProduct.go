func maxProduct(nums []int) int {
    var maxV,secondV int
    for i:=range nums{
        if maxV<=nums[i]{
            secondV=maxV
            maxV=nums[i]
        }else if secondV<nums[i]{
            secondV=nums[i]
        }
    }
    return (maxV-1)*(secondV-1)
}
