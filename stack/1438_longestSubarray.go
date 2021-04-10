func longestSubarray(nums []int, limit int) int {
    /*
    仅需要统计当前窗口内的最大值与最小值，因此我们也可以分别使用两个单调队列解决
    我们使用一个单调递增的队列queMin 维护最小值，一个单调递减的队列queMax 维护最大值。这样我们只需要计算两个队列的队首的差值，即可知道当前窗口是否满足条件
    */
    var queMin,queMAx []int
    var left=0
    var right=0
    var res int
    for right<len(nums){
        // 滑动窗口的最小值，单调递增stack获取最小值
        for len(queMin)>0&&queMin[len(queMin)-1]>nums[right]{
            queMin=queMin[0:len(queMin)-1]//pop 
        }
        queMin=append(queMin,nums[right])
        // 滑动窗口的最大值，单调递减stack获取最大值
        for len(queMAx)>0&&queMAx[len(queMAx)-1]<nums[right]{
            queMAx=queMAx[0:len(queMAx)-1]
        }
        queMAx=append(queMAx,nums[right])
        //不满足条件时收缩左边界
        for abs(queMAx[0]-queMin[0])>limit{
            if nums[left]==queMin[0]{
                queMin=queMin[1:]//pop
            }
            if nums[left]==queMAx[0]{
                queMAx=queMAx[1:]
            }
            left++
        }
        res=max(res,right-left+1)
        right++
    }
    return res
}

func abs(x int)int {
    if x<0 {
        return -1*x
    }
    return x
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
