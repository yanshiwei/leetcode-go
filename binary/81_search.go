func search(nums []int, target int) bool {
    /*
    分为一半后，必定有一部分是有序
    与33题不同，元素可能重复，故严格排序
    如果【left,mid-1】是有序数组， target 的大小满足[left,mid)，则我们应该将搜索范围缩小至 [l, mid - 1]，否则在 [mid + 1, r] 中寻找
    如果【mid,r】是有序数组且 target 的大小满足[mid+1,right]，则我们应该将搜索范围缩小至 [mid + 1, r]，否则在 [l, mid - 1] 中寻找
    因为重复元素，情况3: nums[mid] == nums[left], 无法区分,但是由于nums[mid] != target, 可以left++;
    */
    var left=0;
    var right=len(nums)-1
    var mid int
    for left<=right{
        mid=left+(right-left)/2
        if nums[mid]==target{
            return true
        }
        if nums[left]<nums[mid]{
            //[left mid)有序
            if nums[left]<=target&&nums[mid]>=target {
                //在[left mid]范围
                right=mid-1
            }else {
                //在【mid+1,right】范围
                left=mid+1
            }
        }else if nums[left]>nums[mid]{
            //[mid+1 right]有序
            if nums[mid]<target&&nums[right]>=target {
                //在[mid+1,right]范围内
                left=mid+1
            }else{
                //[left mid]范围
                right=mid-1
            }
        }else {
            //nums[mid] == nums[left], 无法区分,但是由于nums[mid] != target, 可以left++,去掉一个重复干扰项
            left++
        }
    }
    return false
}
