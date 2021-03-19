func movesToMakeZigzag(nums []int) int {
    /*
    我们遍历数组，分两种情况讨论，
    1.要么将奇数位置的元素减少到刚好比相邻的偶数位置元素小;
    2.要么将偶数位置的元素减少到刚好比相邻的奇数位置元素小。
    返回两种情况操作较少的作为答案
    */
    var n=len(nums)
    var res1,res2 int
    for i:=range nums{
        var dis1,dis2 int
        if i%2==1{
            //奇数要小于偶数，i一定从1开始
            if nums[i]>=nums[i-1]{
                dis1=nums[i]-nums[i-1]+1
            }
            if i<n-1&&nums[i]>=nums[i+1]{
                dis2=nums[i]-nums[i+1]+1
            }
            res1+=max(dis1,dis2)
        }else{
            //偶数要小于奇数，i从0开始
            if i>0&&nums[i]>=nums[i-1]{
                dis1=nums[i]-nums[i-1]+1
            }
            if i<n-1&&nums[i]>=nums[i+1]{
                dis2=nums[i]-nums[i+1]+1
            }
            res2+=max(dis1,dis2)
        }
    }
    return min(res1,res2)
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}

func min(x,y int)int {
    if x<y {
        return x
    }
    return y
}
