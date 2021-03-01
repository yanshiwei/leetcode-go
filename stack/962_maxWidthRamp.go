func maxWidthRamp(A []int) int {
    /*
    1 找出最大宽，首先要求i < j，即i尽可能的小，j尽可能的大；很容易想到i需要从左往右，j从右往左遍历数组
    2 第二个条件A[i] <= A[j]，那么我们让A[i]的值尽可能的小, 然后再从右往左遍历数组，挨个和A[i]比较，就可以得到答案，主要的问题就是怎么让A[i]尽可能的小，可以想到利用单调递减栈：
    2.1 正向扫描数组记录严格的单调递减A[i]的下标
    2.2 反向扫描数组比较A[j]与栈顶元素A[s.top()]的值，如果满足条件，则弹出栈顶元素，不断循环直到栈为空
    */
    var stackIndex []int
    for i:=range A {
        if len(stackIndex)==0||A[i]<A[stackIndex[len(stackIndex)-1]] {
            //这里不选相等，是因为相等的情况下肯定是i值更小得到的宽度最大
            stackIndex=append(stackIndex,i)
        }
    }
    var res int
    for i:=len(A)-1;i>=0;i--{
        for len(stackIndex)>0&&A[i]>=A[stackIndex[len(stackIndex)-1]]{
            res=max(res,i-stackIndex[len(stackIndex)-1])
            stackIndex=stackIndex[0:len(stackIndex)-1]//pop
        }
    }
    return res
}

func max(x, y int)int {
    if x<y {
        return y
    }
    return x
}
