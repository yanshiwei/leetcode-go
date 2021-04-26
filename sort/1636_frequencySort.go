func frequencySort(nums []int) []int {
    var fre=make(map[int]int)
    for i:=range nums{
        fre[nums[i]]++
    }
    var newArr=make([]Info,0)
    for k,v:=range fre{
        newArr=append(newArr,Info{Data:k,Cnt:v})
    }
    sort.Slice(newArr,func(i,j int)bool{
        if newArr[i].Cnt!=newArr[j].Cnt{
            return newArr[i].Cnt<newArr[j].Cnt
        }
        return newArr[i].Data>newArr[j].Data
    })
    var res []int
    for i:=range newArr{
        one:=newArr[i]
        for j:=0;j<one.Cnt;j++{
            res=append(res,one.Data)
        }
    }
    return res
}

type Info struct{
    Data int
    Cnt int
}
