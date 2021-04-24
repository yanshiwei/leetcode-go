func mostVisited(n int, rounds []int) []int {
    /*
    由于马拉松全程只会按照同一个方向跑，中间不论跑了多少圈，对所有扇区的经过次数的贡献都是相同的。因此此题的答案仅与起点和终点相关。从起点沿着逆时针方向走到终点的这部分扇区，就是经过次数最多的扇区。
    */
    var lens=len(rounds)
    var start=rounds[0]
    var end=rounds[lens-1]
    var res []int
    if start<=end{
        for i:=start;i<=end;i++{
            res=append(res,i)
        }
    } else{
        //要求走完马拉松全程，故包括start->n及1->end
        // 由于题目要求按扇区大小排序，因此我们要将区间分成两部分
        for i:=1;i<=end;i++{
            res=append(res,i)
        }
        for i:=start;i<=n;i++{
            res=append(res,i)
        }
    }
    return res
}
