func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    res=res[0:0]//必须初始化，否则提交报错
    helper(nums,0,[]int(nil))
    return res
}
var res [][]int
func helper(nums[]int,start int,set []int) {
    res=append(res,append([]int(nil),set...))
    for i:=start;i<len(nums);i++ {
        if i>start&&nums[i]==nums[i-1]{
            continue
        }
        helper(nums,i+1,append(set,nums[i]))
    }
}
