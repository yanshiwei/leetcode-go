func getStrongest(arr []int, k int) []int {
    var m int
    sort.Ints(arr)
    m=arr[(len(arr)-1)/2]
    var handles []Info
    for i:=range arr{
        handles=append(handles,Info{Ori:arr[i],Abs:abs(arr[i]-m)})
    }
    sort.Slice(handles,func(i,j int)bool{
        if handles[i].Abs!=handles[j].Abs{
            return handles[i].Abs>handles[j].Abs
        }
        return handles[i].Ori>handles[j].Ori
    })
    //fmt.Println(handles)
    var res []int
    for i:=range handles{
        res=append(res,handles[i].Ori)
        if len(res)>=k{
            break
        }
    }
    return res
}

func abs(x int)int {
    if x<0 {
        return -1*x
    }
    return x
}

type Info struct {
    Ori int
    Abs int
}
