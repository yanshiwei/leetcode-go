func minPathSum(grid [][]int) int {
    /*
    创建二维数组dp，与原始网格的大小相同，dp[i][j] 表示从左上角出发到 (i,j) 位置的最小路径和。显然，dp[0][0]=grid[0][0]。对于 dp 中的其余元素，通过以下状态转移方程计算元素值:
    第一行，只能向右走故dp[0][j]=dp[0][j-1]+grid[0][j]
    第一列，只能向下走故dp[i][0]=dp[i-1][0]+grid[i][0]
    其他的只能向下或者向右得到， dp[i][j]=min(dp[i-1][j],dp[i][j-1])+grid[i][j]
    */
    var m=len(grid)
    var n=len(grid[0])
    var dp=make([][]int,m)
    for i:=range dp {
        dp[i]=make([]int,n)
    }
    dp[0][0]=grid[0][0]
    for i:=1;i<m;i++{
        //第一列
        dp[i][0]=dp[i-1][0]+grid[i][0]
    }
    for i:=1;i<n;i++{
        //第一行
        dp[0][i]=dp[0][i-1]+grid[0][i]
    }
    for i:=1;i<m;i++{
        for j:=1;j<n;j++ {
            dp[i][j]=min(dp[i-1][j],dp[i][j-1])+grid[i][j]
        }
    }
    return dp[m-1][n-1]
}

func min(x,y int)int {
    if x<y {
        return x
    }
    return y
}
