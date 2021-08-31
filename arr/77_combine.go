var res[][]int
func combine(n int, k int) [][]int {
    res=nil
    if n<1||k>n{
        return res
    }
    helper(n,k,1,[]int(nil))
    return res
}

func helper(n,k,start int,tmp[]int){
    if len(tmp)==k{
        temp:=make([]int,k)
        copy(temp,tmp)
        res=append(res,temp)
        return
    }
    for i:=start;i<=n;i++{
        tmp=append(tmp,i)
        helper(n,k,i+1,tmp)
        tmp=tmp[:len(tmp)-1]
    }
}
