const INT_MAX = int(^int(0)>>1)
const INT_MIN = ^INT_MAX
func nextGreaterElements(nums []int) []int {
    var res[]int
    if len(nums)<1 {
        return res
    }
    for i:=0;i<len(nums);i++ {
        res=append(res,-1)
    }
    numLen:=len(nums)
    for i:=0;i<numLen;i++ {
        var curBig=nums[i]
        var j=i
        var find bool
        for j<numLen {
            if curBig<nums[j] {
                find=true
                curBig=nums[j]
                break
            }
            j=(j+1)%numLen
            if j==i {
                //loop over
                break
            }
        }
        if find {
            res[i]=curBig
        }
    }
    return res
}
