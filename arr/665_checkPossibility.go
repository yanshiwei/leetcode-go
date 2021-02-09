func checkPossibility(nums []int) bool {
    if len(nums)<1{
        return false
    }
    var cnt int
    for i:=0;i<len(nums)-1;i++ {
        x:=nums[i]
        y:=nums[i+1]
        //nums[i]>nums[i+1]
        if x>y {
            cnt++
            if cnt>1{
                return false
            }
            //第一次，更新顺序 满足nums[i-1]<=nums[i]<=nums[i+1]
            //就能满足 非递减的数列
            if i>0&&nums[i-1]>y{
                nums[i+1]=x
            }
        }
    }
    return true
}
