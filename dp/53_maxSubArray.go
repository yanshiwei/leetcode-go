func maxSubArray(nums []int) int {
    /*
    思路：典型的动态规划
    //dp[i] 代表以nums[i]结尾的连续子数组的最大和
1、确定base case：dp[0] = nums[0]
2、确定状态：以当前节点结尾的子序列的和
3、确定选择：每个节点
4、确定状态转移方程：
可以考虑nums[i]单独成为一段还是加入nums[i]+dp[i-1]，这取决于nums[i]+dp[i-1]
与nums[i]的大小，我们希望获得一个比较大的，于是可以写出这样的动态规划转移方程：
dp[i] = nums[i] > dp[i - 1] + nums[i] ? nums[i] : dp[i - 1] + nums[i]
5、再遍历一遍dp数组，找出最大的dp值即可
    */
    if len(nums)<1{
        return 0
    }
    var dp=make([]int,len(nums))
    dp[0]=nums[0]//bad case
    //状态转移方程
    for i:=1;i<len(nums);i++{
        dp[i]=max(nums[i],dp[i-1]+nums[i])
    }
    fmt.Println(dp)
    //找到最大的子序列和
    var res=INTMIN
    for i:=range dp{
        res=max(res,dp[i])
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
