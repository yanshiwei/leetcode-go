func largestRectangleArea(heights []int) int {
    var res int
    if len(heights)<1 {
        return res
    }
    heights=append(heights,0)//hanlde [1] case
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
