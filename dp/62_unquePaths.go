func uniquePaths(m int, n int) int {
    //1由于只能向下或者向右，每个方格到终点的路径数=左侧方格路径数+上面方格路径数
    //2设置一个m*n的矩阵，a[i][j]代表起点到该点的路径数，则对于第一行或者第一列，永远只有向右或者向下一种走法，故一直都是1，
    //其他点都根据第一点原则进行更新
    if m==1||n==1{
        //只有一行或者一列,只有一种走法
        return 1
    }
    var matrix=make([][]int,m)
    for i:=range matrix{
        matrix[i]=make([]int,n)
    }
    for i:=1;i<m;i++{
        matrix[i][0]=1
        for j:=1;j<n;j++{
            matrix[0][j]=1
            matrix[i][j]=matrix[i][j-1]+matrix[i-1][j]
        }
    }
    return matrix[m-1][n-1]
}
