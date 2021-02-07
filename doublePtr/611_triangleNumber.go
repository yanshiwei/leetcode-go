func triangleNumber(nums []int) int {
    /*
    如果无序，任意两边之和大于第三边，任意两边之差小于第三边
    如果有序，边长从小到大为 a、b、c，当且仅当 a + b > c 这三条边能组成
    1 先排序
    2 固定最长的一条边，运用双指针扫描
    3 nums[l] + nums[r] > nums[i]，同时说明 
    nums[l + 1] + nums[r] > nums[i],
     ...,
    nums[r - 1] + nums[r] > nums[i]
        一共r-l个。
    然后right-1缩小大小
    4 如果 nums[l] + nums[r] <= nums[i]，left+1扩大大小
    */
    var m=len(nums)
    var res int
    if m<3{
        return res
    }
    sort.Ints(nums)
    //固定最长的一条边，运用双指针扫描
    //left min is 0
    //right min is 1
    for i:=m-1;i>=2;i--{
        left:=0
        right:=i-1
        for left<right {
            //nums[l] + nums[r] > nums[i]
            if nums[left]+nums[right]>nums[i]{
                res+=right-left
                right--
            }else {
                left++
            }
        }
    }
    return res
}
