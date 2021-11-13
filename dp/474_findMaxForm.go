func findMaxForm(strs []string, m int, n int) int {
    var length=len(strs)
    var dp=make([][][]int,length+1)//dp[i][j][k] 表示输入字符串在子区间 [0, i] 能够使用 j 个 0 和 k 个 1 的字符串的最大数量
    for i:=range dp{
        dp[i]=make([][]int,m+1)
        for j:=range dp[i]{
            dp[i][j]=make([]int,n+1)
        }
    }
    // init is 0 for dp[0]
    for i:=1;i<=length;i++{
        // 第i个字符串
        count:=countZeroAndOne(strs[i-1])//第i个字符串0 1 个数，即当前串的0 1 个数
        for j:=0;j<=m;j++{
            for k:=0;k<=n;k++{
                dp[i][j][k]=dp[i-1][j][k]// default不使用当前串
                zerosCnt:=count[0]
                onesCnt:=count[1]
                if j>=zerosCnt&&k>=onesCnt{
                    //使用当前串
                    dp[i][j][k]=max(dp[i-1][j][k],dp[i-1][j-zerosCnt][k-onesCnt]+1)
                }
            }
        }
    }
    return dp[length][m][n]
}
func countZeroAndOne(str string)[]int{
    var count []int=make([]int,2)
    for i:=range str{
        if str[i]=='0'{
            count[0]++
        }else{
            count[1]++
        }
    }
    return count
}
func max(x,y int)int {
    if x<y{
        return y
    }
    return x
}
