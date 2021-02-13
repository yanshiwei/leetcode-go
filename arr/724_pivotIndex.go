func pivotIndex(nums []int) int {
    /*
    数组总元素和t，当遍历到nums[i]，此时左边连续和sum
    若与右边相等则 sum=t-nums[i]-sum
    即2*sum+nums[i]=t
    */
    var res=-1
    var m=len(nums)
    var t int
    for i:=range nums {
        t+=nums[i]
    }
    var sum int
    for i:=0;i<m;i++{
        if 2*sum+nums[i]==t{
            return i
        }
        sum+=nums[i]
    }
    return res
}
