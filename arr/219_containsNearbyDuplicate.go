func containsNearbyDuplicate(nums []int, k int) bool {
    var m=len(nums)
    for i:=0;i<=m-1;i++{
        end:=i+k
        if end>=m {
            end=m-1
        }
        //fmt.Println(i,end)
        for j:=i+1;j<=end;j++ {
            if nums[i]==nums[j] {
                return true
            }
        }
    }
    return false
}
