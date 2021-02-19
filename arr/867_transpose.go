func transpose(matrix [][]int) [][]int {
    //尺寸为 R x C 的矩阵 A 转置后会得到尺寸为 C x R 的矩阵 ans，对此有 ans[c][r] = A[r][c]
    var res [][]int
    var m=len(matrix)
    if m<1 {
        return res
    }
    var n=len(matrix[0])
    if n<1{
        return res
    }
    res=make([][]int,n)
    for i:=range res{
        res[i]=make([]int,m)
    }
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            res[j][i]=matrix[i][j]
        }
    }
    return res
}
