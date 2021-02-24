func isMonotonic(A []int) bool {
    var m=len(A)
    if m<2 {
        return true
    }
    var isAsc bool//递增
    var last=A[0]
    for i:=1;i<m;i++{
        if last==A[i]{
            continue
        }
        if last<=A[i]{
            isAsc=true
        }
        break
    }

    for i:=1;i<m;i++{
        if isAsc{
            //递增
            if last>A[i]{
                return false
            }
        }else{
            //递减
            if last<A[i]{
                return false
            }
        }
        last=A[i]
    }
    return true
}
