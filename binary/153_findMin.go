func findMin(nums []int) int{
    /*
    二分法
    设左边界left，右边界right ，中间值mid对应的值为pivot。我们将pivot与右边界值对比，由于无重复元素，存在以下两种情况：
    1 pivot小于右边界的值，则[mid,right]都是升序。说明pivot是要寻找的最小值的右边部分的元素；故应该忽略mid的右半部分；
    2 pivot大于右边界的值，则[left,mid]都是升序，说明pivot是要寻找的最小值的左边部分的元素，可以忽略mid左半部分
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
        }else {
            //pivot大于右边界的值，则[left,mid]都是升序,可以忽略mid左半部分
            left=mid+1
        }
    }
    return nums[left]
}
/*
func findMin(nums []int) int {
     将数组从中间分开成左右两部分A B 的时候，A 和B中一定有一部分的数组是有序的。拿示例[4,5,6,7,0,1,2]来看，我们从 6 这个位置分开以后数组变成了 [4, 5, 6] 和 [7, 0, 1, 2] 两个部分，其中左边 [4, 5, 6] 这个部分的数组是有序的。且有序的部分最小值肯定不在此！！。
    根据有序的那个部分确定我们该如何改变二分搜索的上下界，因为我们能够根据有序的那部分判断出 target 在不在这个部分：
    1. 如果 [l, mid - 1] 是有序数组，且 target 的大小满足[left,mid]，则我们应该将搜索范围缩小至 [l, mid - 1]，否则在 [mid + 1, r] 中寻找
    2. 如果 [mid, r] 是有序数组，且 target 的大小满足[mid+1,right]，则我们应该将搜索范围缩小至 [mid + 1, r]，否则在 [l, mid - 1] 中寻找
    //
    var left=0
    var right=len(nums)-1
    if nums[right]>nums[0]{
        return nums[0]
    }
    return helper(nums,left,right)
}

func helper(nums[]int,left,right int)int {
    if left==right{
        return nums[left]
    }
    var mid=left+(right-left)/2
    if nums[mid]>nums[left]{
        //[left,mid]是有序
        //左边序列的最小值就是nums[left]
        //右边序列的最小值是d(nums, mid + 1, right)
        return min(nums[left],helper(nums,mid+1,right))
    }
    //否则[mid+1,right]有序
    return min(nums[mid+1],helper(nums,left,mid))
}
func min(x,y int)int {
    if x<y {
        return x
    }
    return y
}
*/
