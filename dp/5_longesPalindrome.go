func longestPalindrome(s string) string {
    /*
    如果它是回文串，并且长度大于 2，那么将它首尾的两个字母去除之后，它仍然是个回文串。
    就可以用动态规划的方法解决本题。我们用 P(i,j) 表示字符串 s 的第 i到 j个字母组成的串
    则P(i,j)=true表示si-sj子串是回文则要求s[i+1:j−1] 是回文串，并且 s的第 i 和 j 个字母相同时
    对于边界：
    长度为1:P(i,i)=true
    长度为2：P(i,i)=true则要求两个字母相同
    */
    var n=len(s)
    if n<2{
        return s
    }
    var maxLen int=1//至少是一个
    var begin int
    var dp=make([][]bool,n)// dp[i][j] 表示 s[i..j] 是否是回文串
    for i:=range dp{
        dp[i]=make([]bool,n)
        dp[i][i]=true// 初始化：所有长度为 1 的子串都是回文串
    }
    // 先枚举子串长度l
    for l:=2;l<=n;l++{
        // 枚举左边界i
        for i:=0;i<n;i++{
            // 由 L 和 i 可以确定右边界，即 j - i + 1 = l 得
            j:=l+i-1
            // 如果右边界越界
            if j>=n{
                break
            }
            if s[i]!=s[j]{
                dp[i][j]=false
            }else{
                //s的第 i 和 j 个字母相同时
                if j-i<3{
                    dp[i][j]=true
                }else{
                    //长度超过2,要求s[i+1:j−1] 是回文串
                    dp[i][j]=dp[i+1][j-1]
                }
            }
            // 只要 dp[i][L] == true 成立，就表示子串 s[i..L] 是回文，此时记录回文长度和起始位置
            if dp[i][j]&&j-i+1>maxLen{
                maxLen=j-i+1
                begin=i
            }
        }
    }
    return s[begin:begin+maxLen]
}
