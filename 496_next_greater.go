func nextGreaterElement(nums1 []int, nums2 []int) []int {
    var res =make([]int,len(nums1))
    if len(nums1)<1 || len(nums2)<1 {
        return res
    }
    for i:=range res {
        res[i]=-1
    }
    var mapData=make(map[int]int)//because nums1 nums2 is only
    for i,x:=range nums2 {
        mapData[x]=i
    }
    for i,x:=range nums1 {
        if index,ok:=mapData[x];ok {
            for ii:=index+1;ii<len(nums2);ii++ {
                if nums2[ii]>x {
                    res[i]=nums2[ii]
                    break
                }
            }
        }
    }
    return res
}
