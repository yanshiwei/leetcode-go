func checkIfExist(arr []int) bool {
    var fre=make(map[int]int)
    for i:=range arr{
        fre[arr[i]]++
    }
    if v,ok:=fre[0];ok{
        if v%2==0 {
            return true
        }
    }
    for i:=range arr{
        if arr[i]==0 {
            continue
        }
        if _,ok:=fre[arr[i]*2];ok==true{
            return true
        }
    }
    return false
}
