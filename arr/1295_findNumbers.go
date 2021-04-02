func findNumbers(nums []int) int {
    var res int
    for i:=range nums{
        num:=nums[i]
        str:=strconv.Itoa(num)
        if len(str)%2==0{
            res++
        }
    }
    return res
}
