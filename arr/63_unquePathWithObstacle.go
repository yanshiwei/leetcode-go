func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    /*
    与62题类似，只不过要单独处理障碍。
    62题做法是：
    //1由于只能向下或者向右，每个方格到终点的路径数=左侧方格路径数+上面方格路径数
    //2设置一个m*n的矩阵，a[i][j]代表起点到该点的路径数，则对于第一行或者第一列，永远只有向右或者向下一种走法，故一直都是1，
    //其他点都根据第一点原则进行更新
    63题做法是：
    1由于只能向下或者向右，每个方格到终点的路径数=左侧方格路径数+上面方格路径数
    2 设置一个m*n的矩阵，a[i][j]代表起点到该点的路径数，则对于第一行或者第一列，一旦遇到障碍，后续方格就无法到达，故后续路径数目只能0
    //其他点，若当前方格为障碍物，直接赋值0,否则按照第一点更新
    //增加障碍物判断
    */
    var m=len(obstacleGrid)
    var n=len(obstacleGrid[0])
    if obstacleGrid[0][0]==1{
        //起点被阻塞
        return 0
    }
    var matrix=make([][]int,m)
    for i:=range matrix{
        matrix[i]=make([]int,n)
    }
    //初始化第一列
    for i:=0;i<m;i++{
        if obstacleGrid[i][0]==1{
            break
        }
        matrix[i][0]=1
    }
    //初始化第一行
    for i:=0;i<n;i++{
        if obstacleGrid[0][i]==1{
            break
        }
        matrix[0][i]=1
    }
    for i:=1;i<m;i++{
        for j:=1;j<n;j++{
            if obstacleGrid[i][j]==1 {
                matrix[i][j]=0
            }else {
                matrix[i][j]=matrix[i-1][j]+matrix[i][j-1]
            }
        }
    }
    return matrix[m-1][n-1]
}
