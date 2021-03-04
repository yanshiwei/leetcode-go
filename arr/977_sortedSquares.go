func sortedSquares(nums []int) []int {
    var m=len(nums)
    var res[]int
    if m<1 {
        return res
    }
    for i:=0;i<m;i++ {
        tmp:=nums[i]*nums[i]
        res=append(res,tmp)
    }
    sort.Ints(res)
    return res
}
