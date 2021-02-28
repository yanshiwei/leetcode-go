func validMountainArray(arr []int) bool {
    var m=len(arr)
    if m<3 {
        return false
    }
    var maxIdx=0
    var last=arr[0]
    var i int
    for i=1;i<m;i++{
        if last==arr[i]{
            return false
        }
        if arr[maxIdx]<arr[i]{
            maxIdx=i
        }
        last=arr[i]
    }
    last=arr[0]
    var smallCnt,bitCnt int 
    for i1:=1;i1<=maxIdx;i1++{
        if last>=arr[i1]{
            return false
        }
        last=arr[i1]
        smallCnt++
    }
    last=arr[maxIdx]
    for i1:=maxIdx+1;i1<m;i1++{
        if last<=arr[i1]{
            return false
        }
        last=arr[i1]
        bitCnt++
    }
    if smallCnt<1||bitCnt<1 {
        return false
    }
    return true
}
