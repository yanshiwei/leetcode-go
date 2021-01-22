func merge(nums1 []int, m int, nums2 []int, n int)  {
    var oneRealLen=len(nums1)
    var i,j int
    for i=m;i<oneRealLen;i++{
        nums1[i]=nums2[j]
        j++
        if j>=n {
            break
        }
    }
    sort.Ints(nums1)
}
