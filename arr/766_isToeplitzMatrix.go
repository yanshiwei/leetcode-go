func isToeplitzMatrix(matrix [][]int) bool {
    //r1-c1==r2-c2说明同一个对角线
    var fre=make(map[int]int)
    var m=len(matrix)
    if m<1 {
        return false
    }
    var n=len(matrix[0])
    if n<1 {
        return false
    }
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            if v,ok:=fre[i-j];ok==false{
                fre[i-j]=matrix[i][j]
            }else{
                if v!=matrix[i][j]{
                    return false
                }
            }
        }
    }
    return true
}
