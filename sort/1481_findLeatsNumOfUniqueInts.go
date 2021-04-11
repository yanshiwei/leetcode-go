func findLeastNumOfUniqueInts(arr []int, k int) int {
    var fre=make(map[int]int)
    for i:=range arr{
        fre[arr[i]]++
    }
    var ans []Info
    for k,v:=range fre{
        ans=append(ans,Info{Data:k,Fre:v})
    }
    sort.Slice(ans,func(i,j int)bool{
        return ans[i].Fre<ans[j].Fre
    })
    //fmt.Println(ans)
    var n=len(ans)
    for i:=range ans{
        k-=ans[i].Fre
        n--
        if k<=0{
            if k<0 {
                n++
            }
            break
        }
    }
    return n
}

type Info struct{
    Data int
    Fre int
}
