func trimMean(arr []int) float64 {
    var m=len(arr)
    var n=m/20
    sort.Ints(arr)
    var sum float64
    for i:=n;i<m-n;i++{
        sum+=float64(arr[i])
    }
    return sum/float64(m-2*n)
}
