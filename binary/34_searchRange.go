func searchRange(nums []int, target int) []int {
    /*
    使用两个二分查找遍历分别找到最左端和最右端
    使用二分查找的变体，在找到target时并不返回，而是收紧右边界，继续查找，直到找到target的左边界或者找不到则返回它应当出现的位置（即第一个大于target元素的位置）。

再去查找target+1的左边界，减一即为target的右边界。即使target+1不在数组中也没关系，因为总会返回target+1应该所处的位置（即大于target的第一个元素位置）
    */
    if len(nums)<1{
        return []int{-1,-1}
    }
    left:=binarySearchForLeft(nums,target)
    if left<0{
        return []int{-1,-1}
    }
    right:=binarySearchForRight(nums,target)
    return[]int{left,right}
}

func binarySearchForLeft(nums[]int,target int)int {
    var left=0
    var right=len(nums)-1
    var mid int
    for left<=right{
        mid=left+(right-left)/2
        if nums[mid]==target{
            right= mid-1//-1 is left
        }else if nums[mid]<target{
            left=mid+1
        }else {
            right=mid-1
        }
    }
    if left>=len(nums)||nums[left]!=target{
        return -1
    }
    return left
}

func binarySearchForRight(nums[]int,target int)int {
    var left=0
    var right=len(nums)-1
    var mid int
    for left<=right{
        mid=left+(right-left)/2
        if nums[mid]==target{
            left= mid+1//+1 is right
        }else if nums[mid]<target{
            left=mid+1
        }else {
            right=mid-1
        }
    }
    if right<0||nums[right]!=target{
        return -1
    }
    return right
}
