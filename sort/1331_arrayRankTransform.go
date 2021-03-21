func arrayRankTransform(arr []int) []int {
    var copySlice=make([]int,len(arr))
    var fre=make(map[int]int)
    var idx=1
    for i:=range arr{
        copySlice[i]=arr[i]
    }
    sort.Ints(arr)
    for i:=range arr{
        if _,ok:=fre[arr[i]];ok{
            continue
        }
        fre[arr[i]]=idx
        idx++
    }
    for i:=range copySlice{
        v:=fre[copySlice[i]]
        copySlice[i]=v
    }
    return copySlice
}
