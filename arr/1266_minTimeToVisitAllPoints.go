func minTimeToVisitAllPoints(points [][]int) int {
    /*
    访问两个点之间所需要的最小时间仅取决于两个点对应坐标差值的绝对值最大的一个值
    假设两个点坐标分别（x1,y1）和x2,y2，它们差值是dx=abs(x2-x1),dy=abs(y2-y1)
    1.dx==dy, 最小距离就是对角线 dx ，移动耗时dx
    2.dx > dy, 先对角线移动dy，再竖向dx-dy,移动耗时合计就是 dy+dx-dy=dx
    3.dx < dy, 先对角线移动dx，在竖向dy-dx,移动耗时合计就是 dx+dy-dx=dy
    总结就是 max(dx,dy),剩下就是遍历每个节点计算结果就好
    */
    var res int
    var x0=points[0][0]
    var y0=points[0][1]
    for i:=1;i<len(points);i++{
        dx:=abs(points[i][0]-x0)
        dy:=abs(points[i][1]-y0)
        res+=max(dx,dy)
        x0=points[i][0]
        y0=points[i][1]
    }
    return res
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}

func abs(x int)int{
    if x<0 {
        return -1*x
    }
    return x
}
