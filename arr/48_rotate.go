func rotate(matrix [][]int)  {
    var n=len(matrix)
    //水平翻转,上半部分和下半部分交换  new[i][j]=old[n-row-1][col]
    for i:=0;i<n/2;i++{
        //n/2
        for j:=0;j<n;j++{
            matrix[i][j],matrix[n-i-1][j]=matrix[n-i-1][j],matrix[i][j]
        }
    }
    //主对角线翻转，对角线左侧和右侧交换 new[i][j]=new[j][i]
    for i:=0;i<n;i++{
        for j:=0;j<i;j++{
            //j<i
            matrix[i][j],matrix[j][i]=matrix[j][i],matrix[i][j]
        }
    }
}
