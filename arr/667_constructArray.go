func constructArray(n int, k int) []int {
    /*
    构造一个递增序列。前k+1个数构造k个差值不同的数，从k+1之后直接全部递增即可
    ai+1-ai>0
    */
    var nums=make([]int,n)
    nums[0]=1
    var interval=k
    for i:=1;i<=k;i++ {
        if i%2==1{
            //奇数位置
            nums[i]=nums[i-1]+interval
        }else {
            nums[i]=nums[i-1]-interval
        }
        //保持差值不同
        interval--
    }
    //从i+1开始
    for i:=k+1;i<n;i++{
        nums[i]=i+1
    }
    return nums
}
