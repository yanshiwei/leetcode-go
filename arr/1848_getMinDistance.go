func getMinDistance(nums []int, target int, start int) int {
    var res =10001
    abs:=func(x int)int {
        if x<0 {
            return -1*x
        }
        return x
    }
    for i:=range nums{
        if nums[i]==target{
            if res>abs(i-start){
                res=abs(i-start)
            }
        }
    }
    return res
}
