func numSpecial(mat [][]int) int {
    var res int
    var row=make([]int,len(mat))
    var col=make([]int,len(mat[0]))
    for i:=0;i<len(mat);i++{
        for j:=0;j<len(mat[0]);j++{
            if mat[i][j]==1{
                row[i]++
                col[j]++
            }
        }
    }
    for i:=0;i<len(mat);i++{
        for j:=0;j<len(mat[0]);j++{
            if mat[i][j]==1&&row[i]==1&&col[j]==1{
                res++
            }
        }
    }
    return res
}
