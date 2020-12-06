type Info struct {
    Index int
    Data int
}
func dailyTemperatures(T []int) []int {
    var res []int=make([]int,len(T))
    if len(T)<1 {
        return res
    }
    //desc stack right bigger
    var stackIndex []int
    for i:=0;i<len(T);i++ {
        for len(stackIndex)>0&&T[i]>T[stackIndex[len(stackIndex)-1]] {
            headIndex:=stackIndex[len(stackIndex)-1]
            res[headIndex]=i-headIndex
            stackIndex=stackIndex[0:len(stackIndex)-1]//pop
        }
        stackIndex=append(stackIndex,i)
    }
    /*
    var last Info
    for i:=0;i<len(T);i++ {
        last=Info{Index:i,Data:T[i]}
        var j=i
        for j<len(T)&&last.Data>=T[j] {
            j++
        }
        if j>=len(T) {
            res[last.Index]=0
        }else {
            res[last.Index]=j-last.Index
        }
    }
    */
    return res
}
