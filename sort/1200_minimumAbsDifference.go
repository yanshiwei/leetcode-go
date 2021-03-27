func minimumAbsDifference(arr []int) [][]int {
    sort.Ints(arr)
    var res[][]int
    var minDis =INT_MAX
    for i:=0;i<len(arr)-1;i++{
        if minDis>arr[i+1]-arr[i]{
            minDis=arr[i+1]-arr[i]
        }
    }
    for i:=0;i<len(arr)-1;i++{
        if minDis==arr[i+1]-arr[i]{
            res=append(res,[]int{arr[i],arr[i+1]})
        }
    }
    return res
}

const INT_MAX=int(^uint(0)>>1)
