func partitionDisjoint(A []int) int {
    //当A[i] < leftMax时，为了满足left数组的数必须小于等于right中的数，必须将当前A[i]放入left数组，同时更新leftMax和pos，当A[i] = leftMax时暂时没必要将A[i]放入left数组，因为我们求的是最小的left.
    var leftMax=A[0]
    var leftPos=0
    var maxV =A[0]
    for i:=range A{
        maxV=max(maxV,A[i])
        if A[i]>=leftMax{
            continue
        }
        leftMax=maxV
        leftPos=i
    }
    return leftPos+1
}

func max(x,y int)int {
    if x<y {
        return y
    }
    return x
}
