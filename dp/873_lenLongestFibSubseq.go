func lenLongestFibSubseq(arr []int) int {
    /*
    dp
    根据斐波那契数列定义，需要两位而不是一位才能定义一个数列。因此，设dp[i][j]表示以A[i],A[j]结尾且i < j 且 A[i] <A[j]的最长斐波那契式的子序列的长度。
    初始值：
    任意两个元素A[i],A[j]之间最小为2，dp[i][j] = 2
    转移方程：
    当arr[k]+arr[i]=arr[j]那么dp[i][j]=max(dp[i][j],dp[k][i]+1)
    */
    var res int
    var m=len(arr)
    if m<3 {
        return res
    }
    var dp=make([][]int,m-1)//dp[i][j], because i<j and j<m so i<m-1
    for i:=range dp{
        //初始化
        for j:=0;j<m;j++{
            dp[i]=append(dp[i],2)
        }
    }
    var valueMapIndex=make(map[int]int)//arr值及对应的下标
    for j:=0;j<m;j++{
        //数组严格递增
        valueMapIndex[arr[j]]=j
        for i:=0;i<j;i++{
            t:=arr[j]-arr[i]
            if k,ok:=valueMapIndex[t];ok==true&&k<i{
                //说明存在arr[k]+arr[i]=arr[j]且k<i
                dp[i][j]=max(dp[i][j],dp[k][i]+1)
            }
            res=max(res,dp[i][j])
        }
    }
    if res==2{
        //数列肯定大于2
        res=0
    }
    return res
}

func max(x,y int)int{
    if x<y {
        return y
    }
    return x
}
