func findMaxConsecutiveOnes(nums []int) int {
    var m=len(nums)
    var res int
    if m<1 {
        return res
    }
    if m==1 {
        if nums[0]==1 {
            return 1
        }
        return res
    }
    var newSeg bool
    var curLength int
    for i:=0;i<m;i++ {
        if nums[i]==0 {
            newSeg=true
            if curLength>res{
                res=curLength
            }
        }else {
            if newSeg {
                curLength=0
            }
            curLength++
            newSeg=false
        }
    }
    if curLength>res{
        res=curLength
    }   
    return res
}
