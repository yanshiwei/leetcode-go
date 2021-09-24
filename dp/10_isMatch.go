func isMatch(s string, p string) bool {
    var ms=len(s)
    var mp=len(p)
    if ms==0||mp==0{
        return false
    }
    var dp=make([][]bool,ms+1)//dp[i][j]长度j的p串是否匹配长度i的s, seen https://leetcode-cn.com/problems/regular-expression-matching/solution/shou-hui-tu-jie-wo-tai-nan-liao-by-hyj8/
    for i:=range dp{
        dp[i]=make([]bool,mp+1)
        for j:=range dp[i]{
            dp[i][j]=false
        }
    }
    // bad case
    dp[0][0]=true//s p都为空
    for j:=1;j<=mp;j++{
        // s为空串时，但p不为空串，要想匹配，只可能是右端是星号，它干掉一个字符后，把 p 变为空串。
        if p[j-1]=='*'{
            //干掉p[j-2],看p[0,...,j-3]
            dp[0][j]=dp[0][j-2]
        }
    }
    for i:=1;i<=ms;i++{
        for j:=1;j<=mp;j++{
            if s[i-1]==p[j-1]||p[j-1]=='.'{
                //情况1：s[i-1]和 p[j-1] 是匹配的
                dp[i][j]=dp[i-1][j-1]
            }else {
                //s[i-1]和 p[j-1] 是不匹配的
                if p[j-1]!='*'{
                    //如果 p[j-1] 不是星号，那就真的不匹配
                    dp[i][j]=false
                }else{
                    // 如果 p[j-1]是星号
                    if s[i-1]==p[j-2]||p[j-2]=='.'{
                        //s[i-1]和 p[j-2] 是匹配的,分3种情况
                        dp[i][j]=dp[i][j-2]||dp[i-1][j-2]||dp[i-1][j]
                    }else{
                        //s[i-1]和 p[j-2] 是不匹配的
                        dp[i][j]=dp[i][j-2]
                    }
                }
            }
        }
    }
    return dp[ms][mp]
}
