func trap(height []int) int {
    //单调栈，如当前元素大于栈顶，则必然出现凹槽（故单调减栈）
    //计算凹槽长度及高度差即可计算面积
    //高度差计算  当前元素与下一个栈顶元素最小值，然后减去当前栈顶
    if len(height)<3 {
        //at least 3 box
        return 0
    }
    var sum int
    var stackIndex []int
    for i:=range height {
        for len(stackIndex)>0&&height[stackIndex[len(stackIndex)-1]]<height[i] {
            tmp:=height[stackIndex[len(stackIndex)-1]]
            stackIndex=stackIndex[0:len(stackIndex)-1]//pop
            if len(stackIndex)>0 {
                sum+=(min(height[i],height[stackIndex[len(stackIndex)-1]])-tmp)*(i-stackIndex[len(stackIndex)-1]-1)
            }
        }
        stackIndex=append(stackIndex,i)
    }
    return sum
}

func min(x,y int)int {
    if x<y {
        return x
    }
    return y
}
