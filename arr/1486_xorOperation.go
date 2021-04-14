func xorOperation(n int, start int) int {
    var arr=make([]int,n)
    for i:=range arr{
        arr[i]=start+2*i
    }
    var res =arr[0]
    for i:=range arr{
        if i==0 {
            continue
        }
        res^=arr[i]
    }
    return res
}
