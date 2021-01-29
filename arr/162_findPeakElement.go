func findPeakElement(nums []int) int {
    var res int
    if len(nums)<1 {
        return res
    }
    nums=append(nums,INTMIN)
    for i:=1;i<len(nums);i++ {
        if nums[i]>nums[i-1]&&nums[i]>nums[i+1]{
            return i
        }
    }
    return res
}

const INTMAX=int(^uint(0)>>1)
const INTMIN=^INTMAX
