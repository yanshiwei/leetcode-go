func findMin(nums []int) int {
    /*
    二分法
    设左边界left，右边界right ，中间值mid对应的值为pivot。我们将pivot与右边界值对比，由于重复元素，存在以下三种情况：
    1 pivot小于右边界的值，则[mid,right]都是升序。说明pivot是要寻找的最小值的右边部分的元素；故应该忽略mid的右半部分；
    2 pivot大于右边界的值，则[left,mid]都是升序，说明pivot是要寻找的最小值的左边部分的元素，可以忽略mid左半部分
    3 pivot等于右边界值，由于重复元素的存在，我们并不能确定pivot究竟在最小值的左侧还是右侧，因此我们不能莽撞地忽略某一部分的元素。我们唯一可以知道的是，由于它们的值相同，所以无论nums[right] 是不是最小值，都有一个它的「替代品」pivot，因此我们可以忽略二分查找区间的右端点
    */
    var left=0
    var right=len(nums)-1
    var mid int
    for left<right {
        mid=left+(right-left)/2
        pivot:=nums[mid]
        if pivot<nums[right]{
            //pivot小于右边界的值，则[mid,right]都是升序,忽略mid的右半部分
            right=mid
        }else if pivot>nums[right]{
            //pivot大于右边界的值，则[left,mid]都是升序,可以忽略mid左半部分
            left=mid+1
        }else {
            //相等，忽略边界
            right=right-1
        }
    }
    return nums[left]
}
