func findSpecialInteger(arr []int) int {
    var m=len(arr)
    var n=m/4
    var fre=make(map[int]int)
    for i:=range arr{
        fre[arr[i]]++
        if fre[arr[i]]>n{
            return arr[i]
        }
    }
    return 0
}
