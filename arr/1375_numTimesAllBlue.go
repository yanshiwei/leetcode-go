func numTimesAllBlue(light []int) int {
    /*
    遍历数组，记录当前最大亮起来的灯，如果最大亮起来的灯等于遍历过的灯的数量，那么说明前面灯都亮了。举个例子，[2,1,3]，当遍历到3的时候，最大的灯是3，等于当前已经亮起来的灯的个数，因此3左侧的灯全部亮了，算一个可行解。
    */
    var res int
    var curMAx int
    for i:=range light{
        curMAx=max(curMAx,light[i])
        if curMAx==i+1{
            res++
        }
    }
    return res
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
