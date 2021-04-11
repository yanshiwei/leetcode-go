func shuffle(nums []int, n int) []int {
    var out []int
    for i:=0;i<n;i++{
        out=append(out,nums[i]) 
        out=append(out,nums[i+n]) 
    }
    return out
}
