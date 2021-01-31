func missingNumber(nums []int) int {
    var m=len(nums)
    var sum=(m+1)*m/2//sum=n*a1+n(n-1)/2*d
    var oriSum=0
    for i:=range nums {
        oriSum+=nums[i]
    }
    return sum-oriSum
}
