func findDuplicates(nums []int) []int {
    var res[]int
    if len(nums)<1{
        return res
    }
    var m=len(nums)
    for i:=0;i<m;i++ {
        idx:=(nums[i]-1)%m
        //fmt.Println(idx)
        nums[idx]+=m//+m 求余数仍然不变
    }
    fmt.Println(nums)
    for i:=0;i<m;i++{
        if nums[i]-2*m>0 {
            res=append(res,i+1)
        }
    }
    return res
}
