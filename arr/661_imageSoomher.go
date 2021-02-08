func imageSmoother(M [][]int) [][]int {
    var res[][]int
    var m=len(M)
    if m<1 {
        return res
    }
    var n=len(M[0])
    if n<1 {
        return res
    }
    res=make([][]int,m)
    for i:=range res{
        res[i]=make([]int,n)
    }
    for i:=0;i<m*n;i++ {
        curRow:=i/n//第几行
        curCol:=i%n//第几列
        //矩阵改成数组遍历的方式
        //定义每个点的上下左右初始值
        x1:=curCol
        x2:=curCol
        y1:=curRow
        y2:=curRow
        sum:=0
        if curCol>0{
            //非左边界
            x1=curCol-1//确定该点的左边界
        }
        if curCol<n-1{
            //非右边界
            x2=curCol+1//确定该点的右边界
        }
        if curRow>0{
            //非上边界
            y1=curRow-1//确定该点的上边界
        }
        if curRow<m-1{
            //非下边界
            y2=curRow+1//确定该点的下边界 
        }
        //如果条件不满足，该边界就还是他本身
        //遍历
        for j:=y1;j<=y2;j++ {
            for k:=x1;k<=x2;k++ {
                sum+=M[j][k]
            }
        }
        cnt:=(x2-x1+1)*(y2-y1+1)//个数
        res[curRow][curCol]=sum/cnt
    }
    return res
}
