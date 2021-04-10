func kLengthApart(nums []int, k int) bool {
    var m=len(nums)
    if m<k {
        return false
    }
    var lastOneIdx int
    var first =true
    for i:=0;i<m;i++{
        if nums[i]==1{
            if first{
                //
                first=false
            }else{
                if i-lastOneIdx<=k{
                    //fmt.Println(lastOneIdx,i)
                    return false
                }
            }
            lastOneIdx=i
        }
    }
    return true
}
