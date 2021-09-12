func calculateMinimumHP(dungeon [][]int) int {
    //dp,dp[i][j]=max(min(dp[i+1][j],dp[i][j+1])−dungeon(i,j),1).https://leetcode-cn.com/problems/dungeon-game/solution/dong-tai-gui-hua-by-wisemove-2-7/
    var n=len(dungeon)
    var m=len(dungeon[0])
    var dp=make([][]int,n)//dp[i][j]表示从(i,j)到终点所需的最小初始值
    for i:=range dp{
        dp[i]=make([]int,m)
    }
    //init
    dp[n-1][m-1]=max(1-dungeon[n-1][m-1],1)//terminal
    // last col, only to down
    for i:=n-2;i>=0;i--{
        dp[i][m-1]=max(dp[i+1][m-1]-dungeon[i][m-1],1)
    }
    //last row, conly to right
    for j:=m-2;j>=0;j--{
        dp[n-1][j]=max(dp[n-1][j+1]-dungeon[n-1][j],1)
    }
    //inner position
    for i:=n-2;i>=0;i--{
        for j:=m-2;j>=0;j--{
            minv:=min(dp[i+1][j],dp[i][j+1])
            dp[i][j]=max(minv-dungeon[i][j],1)
        }
    }
    return dp[0][0]
}
const INTMAX=int(^uint(0)>>1)
func max(x, y int)int{
    if x<y{
        return y
    }
    return x
}
func min(x,y int)int {
    if x<y {
        return x
    }
    return y
}
