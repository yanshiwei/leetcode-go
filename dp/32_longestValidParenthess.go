func longestValidParentheses(s string) int {
    //dp,dp[i]：以s[i]为结尾的子串中，所形成的最长有效子串的长度.https://leetcode-cn.com/problems/longest-valid-parentheses/solution/shou-hua-tu-jie-zhan-de-xiang-xi-si-lu-by-hyj8/
    var n=len(s)
    var dp=make([]int,n)
    //init is 0
    var maxLen int
    for i:=1;i<n;i++{
        if s[i]==')'{
            if s[i-1]=='('{
                //保底2
                if i-2>=0{
                    //s[i-2]()
                    dp[i]=dp[i-2]+2
                }else{
                    //()
                    dp[i]=2
                }
            }else{
                //s[i-1]和s[i]是'))'，以s[i-1]为结尾的最长有效长度为dp[i-1]，跨过这个长度来看s[i-dp[i-1]-1]这个字符：
                if i-1-dp[i-1]>=0&&s[i-1-dp[i-1]]=='('{
                    if i-2-dp[i-1]>=0{
                        //s[i-dp[i-1]-2]存在,dp[i-2-dp[i-1]]+dp[i-1]+2个）
                        dp[i]=dp[i-1]+2+dp[i-2-dp[i-1]]
                    }else{
                        dp[i]=dp[i-1]+2
                    }
                }else{
                    //s[i-1-dp[i-1]]不存在或等于）
                    dp[i]=0
                }
            }
        }else{
            dp[i]=0
        }
        maxLen=max(maxLen,dp[i])
    }
    return maxLen
}
func max(x, y int)int {
    if x<y {
        return y
    }
    return x
}
