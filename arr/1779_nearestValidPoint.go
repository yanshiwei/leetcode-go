func nearestValidPoint(x int, y int, points [][]int) int {
    var res=-1
    var minDis=20001
    abs:=func(x int)int {
        if x<0 {
            return -1*x
        }
        return x
    }
    for i:=range points{
        p:=points[i]
        if p[0]==x||p[1]==y{
            dis:=abs(p[0]-x)+abs(p[1]-y)
            if minDis>dis{
                minDis=dis
                res=i
            }
        }
    }
    return res
}
