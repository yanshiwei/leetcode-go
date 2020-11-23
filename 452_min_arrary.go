//type Point []int//only two, start end
type PointArr [][]int
func (p PointArr) Len() int {
    return len(p)
}
func (p PointArr) Swap(i,j int) {
    p[i],p[j]=p[j],p[i]
}

func (p PointArr) Less(i,j int) bool {
    return p[i][0]<p[j][0]
}
func min(x,y int) int {
    if x<y {
        return x
    }
    return y
}
func findMinArrowShots(points [][]int) int {
    if len(points)<1 {
        return 0
    }else if len(points)<2 {
        return 1
    }
    //get sort array for point(x,y) arrary
    sort.Sort(PointArr(points))
    var maxEnd=points[0][1]
    //no need start because right space is useless after common space
    //when next point start large than maxEnd thus generate a new common-Point-Task
    var res int=1//at least 1
    for _,one:=range points {
        //one is next point
        if one[0]<=maxEnd {
            maxEnd=min(maxEnd,one[1])//common space
        }else {
            res+=1
            maxEnd=one[1]
        }
    }
    return res
}
