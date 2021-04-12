func findLucky(arr []int) int {
    var res=-1
    var fre=make(map[int]int)
    for i:=range arr{
        fre[arr[i]]++
    }
    var luckArr []Info
    var maxV int
    for k,v:=range fre{
        if k==v {
            if maxV<k{
                maxV=k
            }
            luckArr=append(luckArr,Info{Data:k,Fre:v})
        }
    }
    for i:=range luckArr{
        if maxV==luckArr[i].Data{
            return maxV
        }
    }
    return res
}

type Info struct {
    Data int
    Fre int
}
