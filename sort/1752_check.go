func check(nums []int) bool {
    var sortArr=make([]int,len(nums))
    copy(sortArr,nums)
    sort.Ints(sortArr)
    judgeEqual:=func(delta int)bool{
        for i:=range nums{
            sortIdx:=(i+delta)%len(nums)
            if nums[i]!=sortArr[sortIdx]{
                return false
            }
        }
        return true
    }
    for i:=0;i<len(nums);i++{
        if judgeEqual(i){
            return true
        }
    }
    return false
}
