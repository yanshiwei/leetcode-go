func minOperations(s string) int {
    /*
    dp[0][i+1] 表示让 s.charAt(i) 变为 '0' 需要对 s.substring(0, i+1)操作的次数；
    dp[1][i+1] 表示让 s.charAt(i) 变为 '1' 需要对 s.substring(0, i+1)操作的次数。
    故s[i]=='0':
        dp[0][i+1]=dp[1][i]//s[i]是'0'，此时不需要再变'0'，按照规则，则s[i-1]要求是'1'，故只需要看s[i-1]变成'1'的次数
        dp[1][i+1]=dp[0][i]+1//s[i]是'1',此时需要变成'0'，次数是1，同时按照规则按照规则，则s[i-1]要求是'0'
    类似的
    s[i]=='1'
        dp[0][i+1]=dp[1][i]+1
        dp[1][i+1]=dp[0][i]
    初始化条件是：
    dp[0][0]=0
    dp[1][0]=0
    */
    var dp=make([][]int,2)
    for i:=range dp{
        dp[i]=make([]int,len(s)+1)
    }
    for i:=0;i<len(s);i++{
        if s[i]=='0'{
            dp[0][i+1]=dp[1][i]
            dp[1][i+1]=dp[0][i]+1
        }else{
            dp[0][i+1]=dp[1][i]+1
            dp[1][i+1]=dp[0][i]
        }
    }
    getMin:=func(x,y int)int {
        if x<y {
            return x
        }
        return y
    }
    return getMin(dp[0][len(s)],dp[1][len(s)])
}
