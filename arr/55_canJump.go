func canJump(nums []int) bool {
    /*
    1如果某一个作为 起跳点 的格子可以跳跃的距离是 3，那么表示后面 3 个格子都可以作为 起跳点。
    2可以对每一个能作为 起跳点 的格子都尝试跳一次，把 能跳到最远的距离 不断更新。
    3如果可以一直跳到最后，就成功了。
    */
    if len(nums)<2{
        return true
    }
    var maxPos int 
    for i:=0;i<len(nums);i++{
        if i>maxPos{
            return false
        }
        maxPos=max(maxPos,i+nums[i])
    }
    return true
}
func max(x ,y int)int {
    if x<y {
        return y
    }
    return x
}
