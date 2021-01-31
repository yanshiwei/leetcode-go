func combinationSum3(k int, n int) [][]int {
    var nums=make([]int,9)
    for i:=range nums{
        nums[i]=i+1
    }
    res=res[0:0]//初始化
    helper(nums,k,0,n,0,[]int(nil))
    return res
}

var res[][]int

func helper(nums[]int,k int ,start int,target int,depth int,set []int){
    if depth>k {
        return
    }
    if target==0&&len(set)==k {
        res=append(res,append([]int(nil),set...))
        return
    }
    for i:=start;i<len(nums);i++ {
        helper(nums,k,i+1,target-nums[i],depth+1,append(set,nums[i]))
    }
}
