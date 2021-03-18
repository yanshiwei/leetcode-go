func replaceElements(arr []int) []int {
    var m=len(arr)
    for i:=0;i<m-1;i++{
        nextBig:=getMost(arr[i+1:])
        arr[i]=nextBig
    }
    arr[m-1]=-1
    return arr
}

func getMost(arr []int)int {
    var v=arr[0]
    for i:=range arr{
        if v<arr[i]{
            v=arr[i]
        }
    }
    return v
}
