func maximalSquare(matrix [][]byte) int {
    var res int
    var m=len(matrix)
    if m<1{
        return res
    }
    var n=len(matrix[0])
    if n<1{
        return res
    }
    /*dp(i,j) 表示以(i,j) 为右下角，且只包含 1 的正方形的边长最大值
    /若m[i][j]=='1'，则 dp(i,j) 的值由其上方、左方和左上方的三个相邻位置的 dp 值决定。具体而言若dp(i,j)=4,左侧位置 (i, j - 1)，上方位置 (i - 1, j) 和左上位置 (i - 1, j - 1) 均可以作为一个边长为 4 - 1 = 3 的正方形的右下角股
    dp[i][j-1]>=dp[i][j]-1
    dp[i-1][j]>=dp[i][j]-1
    dp[i-1][j-1]>=dp[i][j]-1
    */
    var dp=make([][]int,m)
    for i:=range dp{
        dp[i]=make([]int,n)
    }
    var maxSide int
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            if matrix[i][j]=='1'{
                if i==0||j==0{
                    // edge only one
                    dp[i][j]=1
                }else{
                    //non-edge get min
                    dp[i][j]=min(min(dp[i][j-1],dp[i-1][j]),dp[i-1][j-1])+1
                }
                if maxSide<dp[i][j]{
                    maxSide=dp[i][j]
                }
            }
        }
    }
    return maxSide*maxSide
}

func min(x,y int)int{
    if x<y{
        return x
    }
    return y
}
