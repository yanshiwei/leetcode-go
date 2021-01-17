func merge(intervals [][]int) [][]int {
    /*
    1 将列表中的区间按照左端点升序排序
    2 如果当前区间的左端点在数组 merged 中最后一个区间的右端点之后，那么它们不会重合，我们可以直接将这个区间加入数组 merged 的末尾;
    3 否则，它们重合，我们需要用当前区间的右端点更新数组 merged 中最后一个区间的右端点，将其置为二者的较大值 无需入队列
    */
    sort.Sort(Interval(intervals))
    fmt.Println(intervals)
    var merge [][]int
    for one:=range intervals{
        if len(merge)<1{
            //空直接插入
            merge=append(merge,intervals[one])
        }else if merge[len(merge)-1][1]<intervals[one][0]{
            //当前区间的左侧大于merge的最后一个的右侧，则肯定不会重合
            merge=append(merge,intervals[one])
        }else{
            //当前区间的左侧在merge的最后一个的右侧，肯定会重合,比较当前区间的右侧与merge的最后一个的右侧
            merge[len(merge)-1][1]=max(merge[len(merge)-1][1],intervals[one][1])
        }
    }
    return merge
}
func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
type Interval [][]int

func (h Interval)Len()int{
    return len(h)
}
func (h Interval)Swap(i, j int){
    h[i],h[j]=h[j],h[i]
}

func (h Interval)Less(i,j int)bool {
    return h[i][0]<h[j][0]
}
