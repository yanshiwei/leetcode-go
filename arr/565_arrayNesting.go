func matrixReshape(nums [][]int, r int, c int) [][]int {
    var m=len(nums)
    if m<1 {
        return nums
    }
    var n=len(nums[0])
    if n<1 {
        return nums
    }
    if m*n!=r*c {
        return nums
    }
    var res [][]int
    var tmp[]int
    for i:=0;i<m;i++ {
        for j:=0;j<n;j++ {
            tmp=append(tmp,nums[i][j])
            if len(tmp)>=c{
                //full col
                res=append(res,append([]int{},tmp...))
                tmp=[]int(nil)
            }
        }
    }
    return res
}
