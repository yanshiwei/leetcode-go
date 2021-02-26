func maxSubarraySumCircular(A []int) int {
    /*
    如果是普通数组，其子数组最大和
    maxDp[i]=max(maxDp[i-1]+a[i],ai)
    res=max(res,maxDp[i])
    循环数组分两种情况讨论：
    0 若数组均为负数 直接返回最大值，不用求和
    1 不循环，上面的做法
    2 循环，包含收尾两端，可先求得中间数据的子数组最小和，然后整个数组的总和减去
    */
    var m=len(A)
    if m==1{
        return A[0]
    }
    var sum =A[0]
    var maxDp=make([]int,m)
    var minDp=make([]int,m)
    maxDp[0]=A[0]
    minDp[0]=A[0]
    var maxV=INT_MIN
    var max1=INT_MIN
    var min2=INT_MAX
    //非循环
    for i:=1;i<m;i++{
        maxV=max(maxV,A[i])

        maxDp[i]=max(maxDp[i-1]+A[i],A[i])
        max1=max(max1,maxDp[i])

        minDp[i]=min(minDp[i-1]+A[i],A[i])
        min2=min(min2,minDp[i])
        
        sum+=A[i]
    }
    if maxV<0{
        return maxV
    }
    return max(max1,sum-min2)
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

const INT_MAX=int(^uint(0)>>1)
const INT_MIN=^INT_MAX
