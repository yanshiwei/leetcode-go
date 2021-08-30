var res [][]int
func combinationSum2(candidates []int, target int) [][]int {
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
func helper(candidates[]int,target,start int,tmp[]int){
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
        // 大剪枝：减去 candidates[i] 小于 0，减去后面的 candidates[i + 1]、candidates[i + 2] 肯定也小于 0，因此用 break
        if target-candidates[i]<0{
            break
        }
        // 小剪枝：同一层相同数值的结点，从第 2 个开始，候选数更少，结果一定发生重复，因此跳过，用 continue
        if i>start&&candidates[i]==candidates[i-1]{
            continue
        }
        tmp=append(tmp,candidates[i])
        helper(candidates,target-candidates[i],i+1,tmp)//不重复使用故i+1
        tmp=tmp[:len(tmp)-1]
    }
}
