func maxSumOfThreeSubarrays(nums []int, k int) []int {
    /*
    用dp[i][j]代表到j元素为止，在有i个子数组的情况下最大的值
    对于dp[i][j]，其最大值是下面两种情况之一
    1 要么在j-1个元素的时候已经找到了i个数组的最大值
    2 要么j-k个元素为止有i-1个数组的最大值，然后加上结束于第j个元素的数组的和形成i个数组的最大值
    前缀和sum[i+1]表示i个元素的连续序列和
    */
    var m=len(nums)
    var sum=make([]int,m+1)
    for i:=0;i<m;i++{
        sum[i+1]=sum[i]+nums[i]
    }
    //初始化第0个子数组元素值为0
    var dp=make([][]int,4)
    for i:=range dp{
        dp[i]=make([]int,m+1)
    }
    for i:=1;i<4;i++{
        for j:=k;j<m+1;j++{
            //从k开始
            dp[i][j]=max(dp[i][j-1],dp[i-1][j-k]+sum[j]-sum[j-k])
        }
    }
    //从第3个数组开始 遍历寻找起始节点
    var maxV=INTMIN
    for j:=1;j<m+1;j++ {
        if dp[3][j]>maxV{
            maxV=dp[3][j]
        }
    }
    var res=make([]int,3)
    for i:=3;i>=1;i-- {
        for j:=1;j<m+1;j++{
            if dp[i][j]==maxV{
                //回溯前一个子数组，并保留j最小的
                res[i-1]=j-k
                maxV-=(sum[j]-sum[j-k])
            }
        }
    }
    return res
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}

const INTMAX=int(^uint(0)>>1)
const INTMIN=^INTMAX
