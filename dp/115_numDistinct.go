
func numDistinct(s string, t string) int {
    var m=len(s)
    var n=len(t)
    if m<n{
        return 0
    }
    var dp=make([][]int,m+1)
    //dp[i][j] s[:i]中存在子序列t[:j]的个数
    for i:=0;i<=m;i++{
        dp[i]=make([]int,n+1)
    }
    // s is empty,任何非空串都不是空串的子序列
    // case
    for i:=0;i<=m;i++{
        //t is empty, 空串是任何串的子序列,
        dp[i][0]=1
    }
    for i:=1;i<=m;i++{
        for j:=1;j<=n;j++{
            if s[i-1]==t[j-1]{
                //s[i-1]==t[j-1]
                // case 1:s[i-1]与t[j-1]也参与匹配，则只需要考虑s[0:i-1]与t[0:j-1];
                // case 2:s[i-1]与t[j-1]不参与匹配，则需要考虑t[0:j]是s[0:i-1]的子序列;
                dp[i][j]=dp[i-1][j-1]+dp[i-1][j]
            }else{
                //s[i-1]!=t[j-1]
                // only case:需要考虑t[0:j]是s[0:i-1]的子序列
                dp[i][j]=dp[i-1][j]
            }
        }
    }
    return dp[m][n]
}
