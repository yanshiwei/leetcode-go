func specialArray(nums []int) int {
    var res=-1
    var idx=0
    for {
        var cnt int
        for i:=range nums{
            if nums[i]>=idx{
                cnt++
            }
        }
        if cnt==idx{
            fmt.Println(cnt,idx)
            return idx
        }
        cnt=0
        idx++
        if idx>100{
            break
        }
    }
    return res
}
