func combinationSum(candidates []int, target int) [][]int {
    /*回溯法
    1 选择。选择一个元素
    2 条件
    3 知道和与target相同
    博客：https://blog.csdn.net/yanerhao/article/details/68561290?utm_source=blogxgwz2
    */
    sort.Ints(candidates)
    res=res[0:0]//必须初始化，否则提交报错
    helper(candidates,0,target,[]int(nil))
    return res
}
var res [][]int
//start 开始遍历
func helper(candidates[]int,start int,target int,set[]int){
    if target<0{
        //不满足
        return
    }
    if target==0{
        //结果结合set 正好和为taregt 
        res=append(res,append([]int(nil),set...))
        //不能res=append(res,set),否则set会被之后值覆盖
        return
    }
    for i:=start;i<len(candidates);i++{
        if i>=1&&candidates[i]==candidates[i-1]{
            //数字可以无限制重复被选取,故可以不考虑不同位置相同的元素
            continue
        }
        helper(candidates,i,target-candidates[i],append(set,candidates[i]))
        //append(set,candidates[i])创建一个新切片 这样递归返回后 仍然是愿切片 不包含candidates[i]相当于pop
    }
    return
}
