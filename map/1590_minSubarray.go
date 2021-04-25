func minSubarray(nums []int, p int) int {
    /*
    假设 nums 的和除以 P，余数是 mod
    1.如果 mod == 0，答案就是 0。
    2.如果 mod != 0，答案变成了找原数组中的最短连续子数组，使得其数字和除以 P，余数也是 mod
    在位置x，前x个数字和对p取模为a，在位置y(y>x)，前y个数字和对p取模为a，那么子区间[x+1,y]的和一定为p的整数倍，因为前后两个余数a相互抵消了。
    在这个题目中，我们也是不断标记前x个数字和并取模，那么如何找到某个子区间的和为对p取余后结果为mod呢，也是一样的，我们对前x个数字和不断取模标记，用sum表示前x个数字和一直被p取模后结果，然后每次找(sum-mod+p)%p的结果是否出现过，如果在位置y找到一样的结果，说明[x+1,y]这个子区间和能对p取模得到mod。
    需要特判一下，如果一开始的数组能全部被p整除，直接返回0，如果一开始的数组和小于p，直接返回-1，不可能。
    */
    var mod int
    var res int
    for i:=range nums{
        mod+=nums[i]
    }
    if mod<p {
        return -1
    }
    mod=mod%p
    if mod==0{
        return res
    }
    var sum int 
    res=1000000001
    var table=make(map[int]int)//余数-位置
    table[0]=-1//余数为0时，当存在一处i的前缀和满足除p后为0，则长度为i+1，故取-1
    //fmt.Println(table)
    for i:=range nums{
        sum+=nums[i]
        cnt:=sum%p
        target:=(cnt-mod+p)%p//+p是为了避免出现负数
        if v,ok:=table[target];ok{
            //fmt.Println(table,i,target)
            res=min(res,i-v)
        }
        table[cnt]=i//将当前位置的前缀和存储map中
    }
    fmt.Println(res)
    if res>=len(nums){
        return -1
    }
    return res
}

func min(x,y int)int {
    if x<y {
        return x
    }
    return y
}
