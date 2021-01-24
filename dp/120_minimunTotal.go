func minimumTotal(triangle [][]int) int {
    /*
    三角形特点：第i行包含i+1个元素
    动态规划
    dp[i][j]代表走到第i行 第j个时最小路径和，i>=0,j>=0,j<=i
    转化公式：
    dp[i][j]=min(dp[i-1][j-1],dp[i-1][j])+t[i][j]
    特殊处理：
    对于第i行最左侧即j=0时，按照题意只能是上一行最左侧即转化公式为
    dp[i][0]=dp[i-1][0]+t[i][0]
    对于第i行最右侧即j=i时，按照题意只能是上一行前一位
    dp[i][i]=dp[i-1][i-1]+t[i][i]
    初始化：
    dp[0][0]=t[0][0]
    最终：
    答案即为 dp[n-1][0]-dp[n-1][n-1]中的最小值，其中 n是三角形的行数
    */
    var m=len(triangle)
    var dp=make([][]int,m)
    dp[0]=[]int{triangle[0][0]}//初始化
    for i:=1;i<m;i++ {
        //第i行长度为i+1,i>=0
        dp[i]=make([]int,i+1)
        //最左侧处理
        dp[i][0]=dp[i-1][0]+triangle[i][0]
        for j:=1;j<len(triangle[i])-1;j++ {
            //len(triangle[i])==i
            dp[i][j]=min(dp[i-1][j-1],dp[i-1][j])+triangle[i][j]
        }
        //最右侧处理
        dp[i][i]=dp[i-1][i-1]+triangle[i][i]
    }
    //fmt.Println(dp)
    return getMin(dp[m-1])
}
func min(x, y int)int {
    if x<y{
        return x
    }
    return y
}

func getMin(x[]int)int {
    var min=x[0]
    for i:=1;i<len(x);i++{
        if min>x[i]{
            min=x[i]
        }
    }
    return min
}
