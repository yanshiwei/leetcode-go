func minSetSize(arr []int) int {
    var m=len(arr)//偶数
    var fre=make(map[int]int)
    var data []Info
    for i:=range arr{
        fre[arr[i]]++
        if fre[arr[i]]>=m/2{
            return 1
        }
    }
    for k,v:=range fre{
        data=append(data,Info{Value:k,Fre:v})
    }
    sort.Slice(data,func(i,j int)bool{
        return data[i].Fre>data[j].Fre
    })
    //fmt.Println(fre,data)
    var res int
    var cnt int
    for i:=range data{
        cnt+=fre[data[i].Value]
        res++
        if cnt>=m/2{
            return res
        }
    }
    return 0
}

type Info struct {
    Value int
    Fre int
}
