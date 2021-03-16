func relativeSortArray(arr1 []int, arr2 []int) []int {
    var res []int
    //sort.Ints()
    for i:=range arr2 {
        ele:=arr2[i]
        for j:=range arr1 {
            if arr1[j]==ele{
                res=append(res,ele)
                arr1[j]=-1
            }
        }
    }
    sort.Ints(arr1)
    var j int
    for j=len(arr1)-1;j>=0;j--{
        if arr1[j]==-1 {
            break
        }
    }
    return append(res,arr1[j+1:]...)
}
