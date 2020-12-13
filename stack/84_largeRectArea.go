func largestRectangleArea(heights []int) int {
    var res int
    if len(heights)<1 {
        return res
    }
    heights=append(heights,0)//hanlde res in stack like[1] case
    //也可以在栈头添加0，因为是当前元素小于栈头才会pop，故这个0永远不会出如此保证栈一定非空这样代码里就不会存在非空栈判断，本代码没用到
    //单增栈，当前元素小于栈顶则pop。每pop一次算一次pop对应面积。以【2，1，5，6，2，3】说明
    //i=0位置，push，第一个元素没法确认最大面积
    //i=1,其高度1比i=0的高度2小pop，i=0没法右扩，故此时可以i=0的最大面积=2*1，i=1还是不确定
    //i=2为5大于i=1的1，则i=2没法确定，因为不确定接下来会不会比i=2的5更大的
    //i=3对应6，也一样没法确定
    //i=4对应2小于i=3的6，此时i=3的高度没法右扩，故此时可以算出i=3的最大面积=6*（4-2-1）=6；pop完i=3，类似地，i=4的2小于i=2的5，同样可以算出i=2的最大面积10*（6-3-1）=10；pop完i=2后，i=4的2大于i=1的1则push
    //i=5的3大于栈顶继续push
    //此时若没有提前家的末尾0，则栈内部还有剩余元素判断，故加了一个0
    //i=6的0，肯定小于栈内所有元素，故可以按照之前逻辑，算出栈内剩余原输入i=5的最大面积=3*1=3；i=4的最大面积2*（6-1-1）=8；i=1的最大面积1*（6）=6
    //注意若栈空了，则此时当前元素最小，则扩宽度是包含当前所有的i
    var st []int//asec stack get first smaller
    //找到右边第一个更小的，此时就可以计算之前单增栈的最大矩形面积
    for i:=range heights {//单调增套路
        for len(st)>0&&heights[i]<=heights[st[len(st)-1]] {
            var headIndex=st[len(st)-1]
            var head=heights[headIndex]
            st=st[0:len(st)-1]//pop//单调增right套路
            //业务逻辑
            var area int
            if len(st)>0 {
                area=head*(i-st[len(st)-1]-1)
            }else{
                //一直都是增，第一个非增
                area=head*i
            }
            res=max(res,area)
        }
        st=append(st,i)//单调增套路
    }
    return res
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
