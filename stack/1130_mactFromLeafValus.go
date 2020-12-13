const IntMax=int(^uint(0)>>1)
//总的原则，大值层次要浅
func mctFromLeafValues(arr []int) int {
    //pattern1 大中小 如3，2，1 从右往左乘
    //pattern2 大小大，6,2,7，按照题意，6和7肯定要往上层传，比较6和7大小。故为了和最小且构造2或者0的子树，6应该比7层次要深，6和2构成一层，7在上一层，
    //pattern3 小中大，左侧增加一个MAX，构造成pattern2
    var stack []int
    var sum int
    stack=append(stack,IntMax)
    for i:=range arr {
        for stack[len(stack)-1]<arr[i] {
            //pattern2
            head:=stack[len(stack)-1]
            stack=stack[0:len(stack)-1]//pop，head决定留下做一个叶子
            sum+=head*min(arr[i],stack[len(stack)-1])//比较左右两边大小
        }
        stack=append(stack,arr[i])
    }
    fmt.Println(sum,stack)
    for len(stack)>2 {
        //has pattern1
        head:=stack[len(stack)-1]
        stack=stack[0:len(stack)-1]//pop
        sum+=head*stack[len(stack)-1]
    }
    return sum
}

func min(x,y int)int {
    if x<y {
        return x
    }
    return  y
}
