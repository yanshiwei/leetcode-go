var res [][]int
func combinationSum(candidates []int, target int) [][]int {
    res=nil
    if len(candidates)<1{
        return res
    }
    sort.Ints(candidates)
    if candidates[0]>target{
        return res
    }
    helper(candidates,target,0,[]int(nil))
    return res
}
func helper(candidates []int,target,start int,tmp []int){
    if target==0{
        temp:=make([]int,len(tmp))
        copy(temp,tmp)
        res=append(res,temp)
        return
    }
    if target<0{
        return
    }
    for i:=start;i<len(candidates);i++{
        //无需保证唯一性，因为数字可以反复
        tmp=append(tmp,candidates[i])
        helper(candidates,target-candidates[i],i,tmp)
        tmp=tmp[:len(tmp)-1]
    }
}
