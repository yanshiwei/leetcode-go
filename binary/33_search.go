func search(nums []int, target int) int {
    /*
    将数组从中间分开成左右两部分A B 的时候，A 和B中一定有一部分的数组是有序的。拿示例[4,5,6,7,0,1,2]来看，我们从 6 这个位置分开以后数组变成了 [4, 5, 6] 和 [7, 0, 1, 2] 两个部分，其中左边 [4, 5, 6] 这个部分的数组是有序的。
    根据有序的那个部分确定我们该如何改变二分搜索的上下界，因为我们能够根据有序的那部分判断出 target 在不在这个部分：
    1. 如果 [l, mid - 1] 是有序数组，且 target 的大小满足[left,mid]，则我们应该将搜索范围缩小至 [l, mid - 1]，否则在 [mid + 1, r] 中寻找
    2. 如果 [mid, r] 是有序数组，且 target 的大小满足[mid+1,right]，则我们应该将搜索范围缩小至 [mid + 1, r]，否则在 [l, mid - 1] 中寻找
    */
    if len(nums)<1{
        return -1
    }
    var left int
    var right=len(nums)-1
    var mid int
    for left<=right{
        mid=(left+right)/2
        if nums[mid]==target{
            return mid
        }
        if nums[left]<=nums[mid] {
            //[left,mid]有序
            if nums[left]<=target&&nums[mid]>=target {
                //在[left,mid]范围内
                right=mid-1
            }else {
                //否则在[mid+1,right]
                left=mid+1
            }
        }else {
            //[mid+1,right]有序
            if nums[mid]<target&&nums[right]>=target {
                //在[mid+1,right]范围内
                left=mid+1
            }else {
                //否则在[left,mid]范围内
                right=mid-1
            }
        }
    }
    return -1
}
