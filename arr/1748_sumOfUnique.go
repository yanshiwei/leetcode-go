func sumOfUnique(nums []int) int {
    var fre=make([]int,101)
    for i:=range nums{
        fre[nums[i]]++
    }
    var sum int
    for k,v:=range fre{
        if v==1{
            sum+=k
        }
    }
    return sum
}
