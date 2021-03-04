func maxTurbulenceSize(arr []int) int {
    /*
    dp
    dp[i][0]代表arr[i]结尾的，且满足arr[i-1]>arr[i]的子数组的最大长度
    dp[i][1]代表arr[i]结尾的，且满足arr[i-1]<arr[i]的子数组的最大长度
    初始化
    dp[0][0]=1,dp[0][1]=0
    转移方程
    1 当arr[i-1]>arr[i]，
        dp[i][0]=dp[i-1][1]+1;
        dp[i][1]=1
    2 当arr[i-1]<arr[i]
        dp[i][0]=1
        dp[i][1]=dp[i-1][0]+1
    3 当arr[i-1]==arr[i],arr[i−1] 和 arr[i] 不能同时出现在同一个湍流子数组
        dp[i][0]=1
        dp[i][1]=1
    */
    var m=len(arr)
    var dp=make([][]int,m)
    //init
    for i:=range dp {
        dp[i]=make([]int,2)
    }
    dp[0][0]=1
    dp[0][1]=1
    for i:=1;i<m;i++{
        if arr[i-1]>arr[i]{
            dp[i][0]=dp[i-1][1]+1
            dp[i][1]=1
        }else if arr[i-1]<arr[i]{
            dp[i][0]=1
            dp[i][1]=dp[i-1][0]+1
        }else {
            dp[i][0]=1
            dp[i][1]=1
        }
    }
    var res int
    for i:=range dp{
        res=max(res,dp[i][0])
        res=max(res,dp[i][1])
    }
    return res
}

func max(x, y int)int {
    if x<y {
        return y
    }
    return x
}
