func insert(intervals [][]int, newInterval []int) [][]int {
    /*
    对于区间s1,s2 如果没有交集，则
    s1在s2左侧，此时r1<l2
    s1在s2右侧，此时l1>r2
    如果重合，交集为
    min(l1,l2),max(r1,r2)
    */
    if len(intervals)<1 {
        intervals=append(intervals,newInterval)
        return intervals
    }
    var merge [][]int
    var left=newInterval[0]
    var right=newInterval[1]
    var insert bool
    for _,interval:=range intervals{
        if interval[0]>right{
            //interval在要插入的区间int[]{left,right}的右侧，没有交集，由于newInterval在interval左侧，故剩下的interval不会再遇到newInterval了，更不会重合了，直接插入
            if insert==false {
                merge=append(merge,[]int{left,right})
                insert=true
            } 
            merge=append(merge,interval)
        }else if interval[1]<left{
            //interval在要插入的区间int[]{left,right}的左侧，没有交集
            //newInterval由于不清楚与后面区间关系，故newInterval暂时不确定是否插入
            merge=append(merge,interval)
        }else{
            //有重合 更新范围
            left=min(left,interval[0])
            right=max(right,interval[1])
        }
    }
    if insert==false{
        //如果不存在这样的区间，我们需要在遍历结束后，将int[]{left,right} 加入答案
        merge=append(merge,[]int{left,right})
    }
    return merge
}

func min(x,y int)int {
    if x<y {
        return x
    }
    return y
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
