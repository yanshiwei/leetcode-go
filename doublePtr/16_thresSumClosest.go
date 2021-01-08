func threeSumClosest(nums []int, target int) int {
    var res int
    if len(nums)<3{
        return res
    }
    sort.Ints(nums)
    res=1e9
    for i:=range nums{
        if i>0&&nums[i]==nums[i-1]{
            continue
            //避免重复
        }
        left:=i+1
        right:=len(nums)-1
        for left<right{
            sum:=nums[left]+nums[right]
            if abs(target,sum+nums[i])<abs(target,res){
                res=sum+nums[i]
            }
            if sum>target-nums[i]{
                right--
                for left<right&&nums[right]==nums[right+1]{right--}
            }else if sum<target-nums[i]{
                left++
                for left<right&&nums[left]==nums[left-1]{left++}
            }else {
                return target
            }
        }
    }
    return res
}
func abs(x,y int)int{
    z:=x-y
    if z>=0{
        return z
    }
    return -z
}
func min(x,y int)int {
    if x<y{
        return x
    }
    return y
}
