const INT_MAX = int(^uint(0) >> 1)//first is 0 res is 1
const INT_MIN = ^INT_MAX//first 1 res is 0
func find132pattern(nums []int) bool {
    var exist bool
    if len(nums)<3 {
        return exist
    }
    var stack []int//firstBig
    var secondBig=INT_MIN
    for i:=len(nums)-1;i>=0;i-- {
        if nums[i]<secondBig {
            return true
        }
        for len(stack)>0&&nums[i]>stack[len(stack)-1] {
            //fron tail to loop so that second big is tail
            //find first big til and get second big
            secondBig=stack[len(stack)-1]
            stack=stack[0:len(stack)-1]//pop
        }
        stack=append(stack,nums[i])
    }
    return exist
}
