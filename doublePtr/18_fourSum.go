func fourSum(nums []int, target int) [][]int {
    /*
    与15的三数和类似（https://leetcode-cn.com/problems/3sum/），排序+双重指针。这里多一重循环而已
    */
    var res [][]int
    if len(nums)<4{
        return res
    }
    sort.Ints(nums)
    for i:=0;i<len(nums)-3;i++{
        //右边必须3个，
        if i>0&&nums[i]==nums[i-1]{
            //避免nums[i]重复
            continue
        }
        for j:=i+1;j<len(nums)-2;j++ {
            //右边只需要有2个
            if j>i+1&&nums[j]==nums[j-1]{
                continue
            }
            //开启双指针
            left:=j+1
            right:=len(nums)-1
            for left<right{
                if nums[i]+nums[j]+nums[left]+nums[right]<target{
                    //太小了
                    left++
                }else if nums[i]+nums[j]+nums[left]+nums[right]>target{
                    //太大了
                    right--
                }else {
                    res=append(res,[]int{nums[i],nums[j],nums[left],nums[right]})
                    left++
				    right--
                    for left<right&&nums[left]==nums[left-1]{left++}
                    for left<right&&nums[right]==nums[right+1]{right--}
                }
            }
        }
    }
    return res
}
