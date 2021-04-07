func createTargetArray(nums []int, index []int) []int {
    var m=len(nums)
    var res []int
    for i:=0;i<m;i++{
        if index[i]<len(res){
            //
            res=append(res,nums[i])
            copy(res[index[i]+1:],res[index[i]:len(res)-1])
            res[index[i]]=nums[i]
        }else{
            res=append(res,nums[i])
        }
        //fmt.Println(res)
    }
    return res
}
