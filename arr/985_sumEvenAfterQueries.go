func sumEvenAfterQueries(A []int, queries [][]int) []int {
    var res []int
    var lastSum int
    for i:=range A {
        if A[i]%2==0{
            lastSum+=A[i]
        }
    }
    for i:=range queries{
        val:=queries[i][0]
        index:=queries[i][1]
        if A[index]%2==0{
            lastSum-=A[index]
        }
        A[index]+=val
        if A[index]%2==0{
            lastSum+=A[index]
        }
        res=append(res,lastSum)
    }
    return res
}
