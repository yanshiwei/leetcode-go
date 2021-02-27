func minFlipsMonoIncr(S string) int {
    /*
    dp[i][0]，dp[i][1] 分别代表字符 S[i] 最终选择 0 和 1 的最少翻转次数
    注意考虑到递增，dp[i][0]只能由dp[i-1][0]转化
    故
    如S[i]='1'
    dp[i][0]=dp[i-1][0]+1// 只能从 0 转化来，翻转 '1' 为 '0'，翻转次数加 1
    dp[i][1]=min(dp[i-1][0],dp[i-1][1])//S[i]已经是1 无需翻转
    如S[i]='0'
    dp[i][0]=dp[i-1][0]//只能从 0 转化来，已经是0无需翻转
    dp[i][1]=min(dp[i-1][0]+1,dp[i-1][1]+1)//翻转 '0' 为 '1'，翻转次数加 1
    初始化
    如S[0]='1'
    dp[0][0]=1
    dp[0][1]=0
    如S[0]='0'
    dp[0][0]=0
    dp[0][1]=1
    */
    var m=len(S)
    var dp=make([][]int,m)
    for i:=range dp{
        dp[i]=make([]int,2)
    }
    //初始化
    if S[0]=='0'{
        dp[0][1]=1
    }else {
        dp[0][0]=1
    }
    //dp
    for i:=1;i<m;i++{
        if S[i]=='1'{
            dp[i][0]=dp[i-1][0]+1
            dp[i][1]=min(dp[i-1][0],dp[i-1][1])
        }else{
            dp[i][0]=dp[i-1][0]
            dp[i][1]=min(dp[i-1][0],dp[i-1][1])+1
        }
    }
    return min(dp[m-1][0],dp[m-1][1])
}

func min(x, y int)int {
    if x<y {
        return x
    }
    return y
}
