func longestSubarray(nums []int) int {
    var left,right int//窗口内只允许1个0
    var maxLen int//最长非空子数组长
    var count int//窗口内0数目
    for right<len(nums){
        if nums[right]==0{
            count++
        }
        for count>1{
            //窗口内0数目超过1
            if nums[left]==0{
                count--
            }
            left++
        }
        //因为题目要求删除一个元素，所以长度计算公式为（end-start）而不是之前的（end-start+1）
        maxLen=max(maxLen,right-left)
        right++
    }
    return maxLen
}

func max(x, y int)int {
    if x<y {
        return y
    }
    return x
}
