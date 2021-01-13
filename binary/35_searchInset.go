func searchInsert(nums []int, target int) int {
    if target<nums[0]{
        nums=append(nums,target)
        nums[0],nums[len(nums)-1]=nums[len(nums)-1],nums[0]
        return 0
    }
    if target>nums[len(nums)-1]{
        nums=append(nums,target)
        return len(nums)-1
    }
    var left=0
    var right=len(nums)-1
    var mid,res int
    for left<=right{
        mid=(left+right)/2
        if nums[mid]==target{
            return mid
        }
        if nums[mid]>target{
            //在[left，mid]范围
            res=mid
            right=mid-1
        }else {
            //[mid+1, right]范围
            left=mid+1
        }
    }
    return res
}
