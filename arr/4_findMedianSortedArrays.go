func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    l1:=len(nums1)
    l2:=len(nums2)
    var nums=make([]int,l1+l2)
    if l1==0 {
        if l2%2==0 {
            //偶数
            return (float64(nums2[l2/2-1])+float64(nums2[l2/2]))/2
        }else {
            //奇数
            return float64(nums2[l2/2])
        }
    }
    if l2==0 {
        if l1%2==0 {
            //偶数
            return (float64(nums1[l1/2-1])+float64(nums1[l1/2]))/2
        }else {
            //奇数
            return float64(nums1[l1/2])
        }
    }
    var i,j,k int
    for i<l1&&j<l2{
        if nums1[i]<nums2[j]{
            nums[k]=nums1[i]
            i++
        }else {
            nums[k]=nums2[j]
            j++
        }
        k++
    }
    for i<l1 {
            nums[k]=nums1[i]
            i++
            k++
    }
    for j<l2 {
            nums[k]=nums2[j]
            j++
            k++
    }
    //fmt.Println(nums)
    if len(nums)%2==0 {
        //偶数
        return (float64(nums[len(nums)/2-1])+float64(nums[len(nums)/2]))/2
    }else {
        //奇数
        return float64(nums[len(nums)/2])
    }
}
