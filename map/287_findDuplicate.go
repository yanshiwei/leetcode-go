func findDuplicate(nums []int) int {
    var m=len(nums)
    if m<1 {
        return 0
    }
    var fre=make(map[int]int)
    for i:=range nums {
        if _,ok:=fre[nums[i]];ok==false {
            fre[nums[i]]=1
        }else {
            return nums[i]
        }
    }
    return 0
}
