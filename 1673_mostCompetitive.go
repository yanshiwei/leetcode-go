func mostCompetitive(nums []int, k int) []int {
    //最小的长度K的连续序列
    //seen https://www.youtube.com/watch?v=CSBTbK-egZo
    var stack []int//asc stack
    var count=len(nums)-k//most can delete
    for i:=range nums {
        for len(stack)>0&&count>0&&stack[len(stack)-1]>nums[i] {
            stack=stack[0:len(stack)-1]//pop
            count--
        }
        stack=append(stack,nums[i])
    }
    for count>0 {
        //like nums = [2,4,3,3,5,4,9,6], k = 4
        stack=stack[0:len(stack)-1]//pop tail is asc
        count--
    }
    return stack
}
