func sortArrayByParity(A []int) []int {
    var res []int
    for i:=range A {
        if A[i]%2==0 {
            res=append(res,A[i])
        }
    }
    for i:=range A {
        if A[i]%2!=0 {
            res=append(res,A[i])
        }
    }
    return res
}
