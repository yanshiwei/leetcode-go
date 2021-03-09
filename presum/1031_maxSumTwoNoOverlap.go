func maxSumTwoNoOverlap(A []int, L int, M int) int {
    /*
     数组可依据索引j划分为两部分，L在j左边，M在j右边或L在j右边，M在j左边。
    # 对于每个j，按照两种情况计算最大的L和M和，并更新答案
    前缀和数组保存前 n 位的和，presum[1]保存的就是 nums 数组中前 1 位的和，也就是 presum[1] = nums[0], presum[2] = nums[0] + nums[1] = presum[1] + nums[1]. 依>次类推
    */
    var m=len(A)
    var preSum=make([]int,m+1)
    //前缀和数组保存前 n 位的和，presum[1]保存的就是 nums 数组中前 1 位的和，也就是 presum[1] = nums[0], presum[2] = nums[0] + nums[1] = presum[1] + nums[1]. 依>次类推
    for i:=1;i<m+1;i++{
        preSum[i]=preSum[i-1]+A[i-1]
    }
    var res=preSum[L+M]//前L+M长序列和
    var lMax=preSum[L]//L子数组和最大值
    var mMax=preSum[M]//M子数组和最大值
    //j代表当前位于右边的数组的末尾索引
    for j:=L+M+1;j<m+1;j++{
        //当L在M前时，j代表M的最后一个索引,此时M已确定
        lMax=max2(lMax,preSum[j-M]-preSum[j-M-L])
        curM:=preSum[j]-preSum[j-M]
        resLM:=lMax+curM
        //当L在M后时，i代表L的最后一个索引，此时L已确定
        mMax=max2(mMax,preSum[j-L]-preSum[j-L-M])
        curL:=preSum[j]-preSum[j-L]
        resML:=mMax+curL
        res=max3(res,resLM,resML)
    }
    return res
}

func max2(x,y int)int{
    if x<y{
        return y
    }
    return x
}

func max3(x,y,z int)int {
    a:=max2(x,y)
    return max2(a,z)
}
