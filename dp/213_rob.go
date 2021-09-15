func rob(nums []int) int {
    var res int
    if len(nums)<1{
        return res
    }
    if len(nums)==1{
        return nums[0]
    }
    if len(nums)==2{
        return max(nums[0],nums[1])
    }
    //假设数组 nums 的长度为 n。如果不偷窃最后一间房屋，则偷窃房屋的下标范围是[0,n−2]；如果不偷窃第一间房屋，则偷窃房屋的下标范围是[1,n−1]。在确定偷窃房屋的下标范围之后，即可用第 198 题的方法解决。对于两段下标范围分别计算可以偷窃到的最高总金额，其中的最大值即为在 n 间房屋中可以偷窃到的最高总金额。
    return max(robInner(nums,0,len(nums)-1),robInner(nums,1,len(nums)))
}

func robInner(nums[]int,l,r int)int{
    var length=r-l
    if length==1{
        return nums[l]
    }
    var dp=make([]int,length)//dp[i]第i个house所获得金额
    dp[0]=nums[l]
    dp[1]=max(nums[l],nums[l+1])
    for i:=2;i<length;i++{
        dp[i]=max(dp[i-1],dp[i-2]+nums[l+i])
    }
    return dp[length-1]
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
