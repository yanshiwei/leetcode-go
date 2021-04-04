func numOfSubarrays(arr []int, k int, threshold int) int {
    var res int
    var m=len(arr)
    if m<k{
        return res
    }
    totoal:=k*threshold
    //初始
    var sum int
    for i:=0;i<k;i++{
        sum+=arr[i]
    }
    if sum>=totoal{
        res++
    }
    for i:=k;i<m;i++{
        sum-=arr[i-k]
        sum+=arr[i]
        if sum>=totoal{
            res++
        }
    }
    return res
}
