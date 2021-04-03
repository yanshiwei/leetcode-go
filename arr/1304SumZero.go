func sumZero(n int) []int {
    if n==1 {
        return []int{0}
    }
    var res[]int
    for i:=1;i<=n/2;i++{
        res=append(res,i)
        res=append(res,-1*i)
    }
    if n%2==1{
        res=append(res,0)
    }
    return res
}
