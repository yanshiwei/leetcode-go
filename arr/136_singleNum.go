func singleNumber(nums []int) int {
    //数组中的全部元素的异或运算结果即为数组中只出现一次的数字
    var res int
    for i:=range nums{
        res^=nums[i]
    }
    return res
}
