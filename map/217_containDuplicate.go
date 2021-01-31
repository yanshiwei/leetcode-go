func containsDuplicate(nums []int) bool {
    var fre=make(map[int]bool)
    for i:=range nums {
        if _,ok:=fre[nums[i]];ok==false {
            fre[nums[i]]=true
        }else {
            return true
        }
    }
    return false
}
