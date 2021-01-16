func combinationSum2(candidates []int, target int) [][]int {
    //https://leetcode-cn.com/problems/combination-sum-ii/solution/hui-su-jian-zhi-by-da-za-cao/
    sort.Ints(candidates)
    res=res[0:0]//必须要清理，否则提交报错
    helper(candidates,0,target,[]int(nil))
    return res
}
var res [][]int
func helper(candidates[]int,start int,target int,set []int){
    //fmt.Println(start,target,set)
    if target<0{
        return
    }
    if target==0 {
        res=append(res,append([]int(nil),set...))
        //不能res=append(res,set),否则set会被之后值覆盖
        return
    }
    for i:=start;i<len(candidates);i++{
        if i > start && candidates[i] == candidates[i-1] {
            continue
        }
        //每个数字在每个组合中只能使用一次 故不需要去重
        helper(candidates,i+1,target-candidates[i],append(set,candidates[i]))
    }
    return
}
