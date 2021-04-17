func rangeSum(nums []int, n int, left int, right int) int {
    var sum=make([]int,n*(n+1)/2)
    var idx int
    for i:=0;i<n;i++{
        var tmpSum int
        for j:=i;j<n;j++{
            tmpSum+=nums[j]
            sum[idx]=tmpSum
            idx++
        }
    }
    //fmt.Println(sum)
    sort.Ints(sum)
    var res int
    for i:=left-1;i<right;i++{
        res=(res+sum[i])%1000000007
    }
    return res
}
