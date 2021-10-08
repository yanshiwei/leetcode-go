func isMatch(s string, p string) bool {
    //dp[i][j]表示字符串 ss 的前 ii 个字符和模式 pp 的前 jj 个字符是否能匹配
    var ms=len(s)
    var mp=len(p)
    var dp=make([][]bool,ms+1)
    for i:=range dp{
        dp[i]=make([]bool,mp+1)
    }
    //init
    dp[0][0]=true//当字符串 ss 和模式 pp 均为空时，匹配成功；
    for i:=1;i<=ms;i++{
        dp[i][0]=false//空模式无法匹配非空字符串；
    }
    //非空模式匹配空字符串
    for j:=1;j<=mp;j++{
        if p[j-1]=='*'{
            //星号才能匹配空字符串
            dp[0][j]=true
        }else{
            //只有当模式 p 的前 j个字符均为星号时，dp[0][j] 才为真
            break
        }
    }
    for i:=1;i<=ms;i++{
        for j:=1;j<=mp;j++{
            if p[j-1]==s[i-1]{
                //p[j-1]is letter
                dp[i][j]=dp[i-1][j-1]
            }else if p[j-1]=='?'{
                //p[j-1] is ?
                dp[i][j]=dp[i-1][j-1]
            }else if p[j-1]=='*'{
                //p[j-1]is *, not use and use *
                dp[i][j]=dp[i-1][j]||dp[i][j-1]
            }
        }
    }
    return dp[ms][mp]
}
