func containsPattern(arr []int, m int, k int) bool {
    if len(arr)<m*k {
        return false
    }
    var i,j int
    for i=0;i<=len(arr)-m*k;i++{
        for j=i+m;j<i+m*k;j++{
            if arr[j]!=arr[j-m]{
                break
            }
        }
        if j==i+m*k {
            return true
        }
    }
    return false
}
