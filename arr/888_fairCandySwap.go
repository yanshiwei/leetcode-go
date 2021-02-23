func fairCandySwap(A []int, B []int) []int {
    var res[]int
    var ma=len(A)
    if ma<1{
        return res
    }
    var mb=len(B)
    if mb<1{
        return res
    }
    var sum,sumA int
    for i:=range A {
        sum+=A[i]
    }
    sumA=sum
    var Bmap=make(map[int]int)
    for i:=range B {
        sum+=B[i]
        Bmap[B[i]]=B[i]
    }
    for i:=range A {
        t:=sum/2-(sumA-A[i])
        if k,ok:=Bmap[t];ok==true{
            res=append(res,[]int{A[i],k}...)
            return res
        }
    }
    return res
}
