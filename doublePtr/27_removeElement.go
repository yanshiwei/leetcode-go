func removeElement(nums []int, val int) int {
    if len(nums)<1{
        return 0
    }
    var i int
    var j int
    for j=0;j<len(nums);j++{
        if val==nums[j]{
            continue
        }else {
            nums[i]=nums[j]
            i++
        }
    }
    return i
}
