func setZeroes(matrix [][]int)  {
    var rowMap=make(map[int]bool)
    var colMap=make(map[int]bool)
    for i:=0;i<len(matrix);i++{
        for j:=0;j<len(matrix[0]);j++{
            if matrix[i][j]==0 {
                if _,ok:=rowMap[i];ok==false{
                    rowMap[i]=true
                }
                if _,ok:=colMap[j];ok==false{
                    colMap[j]=true
                }
            }
        }
    }
    for i:=0;i<len(matrix);i++{
        for j:=0;j<len(matrix[0]);j++{
            if _,ok:=rowMap[i];ok{
                matrix[i][j]=0
            }
            if _,ok:=colMap[j];ok{
                matrix[i][j]=0
            }
        }
    }
}
