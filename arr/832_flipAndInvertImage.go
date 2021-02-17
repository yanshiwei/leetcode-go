func flipAndInvertImage(A [][]int) [][]int {
    var m=len(A)
    if m<1 {
        return [][]int{}
    }
    var n=len(A[0])
    if n<1 {
        return [][]int{}
    }
    for i:=0;i<m;i++{
        var tmp []int
        for j:=n-1;j>=0;j--{
            tmp=append(tmp,A[i][j])
        }
        A[i]=tmp
    }
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            if A[i][j]==0{
                A[i][j]=1
            }else {
                A[i][j]=0
            }
        }
    }
    return A
}
