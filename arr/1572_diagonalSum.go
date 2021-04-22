func diagonalSum(mat [][]int) int {
    var res int
    for i:=0;i<len(mat);i++{
        for j:=0;j<len(mat[0]);j++{
            if i-j==0{
                res+=mat[i][j]
            }
            if i+j==len(mat)-1{
                res+=mat[i][j]
            }
            if i-j==0&&i+j==len(mat)-1{
                fmt.Println(mat[i][j])
                res-=mat[i][j]
            }
        }
    }
    return res
}
