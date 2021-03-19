func decompressRLElist(nums []int) []int {
    var res []int
    var m=len(nums)
    //var fre=make(map[int]int)
    for i:=0;i<m;i++{
        if i%2==1{
            //fmt.Println(i,nums[i-1])
            for j:=0;j<nums[i-1];j++{
                res=append(res,nums[i])
            }
        }
    }
    return res
}
