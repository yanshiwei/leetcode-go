func sumSubarrayMins(arr []int) int {
    //seen in https://blog.csdn.net/qq_17550379/article/details/86548585
    const MOD=1e9+7
    var res int
    if len(arr)<1 {
        return res
    }
    var stackIndex []int//asec stack
    var before=make([]int,len(arr))//num
    var after=make([]int,len(arr))//num
    //min beofore delta1 after delta2 so the arr num is (delta1+1)*(delata2+1)
    //find the first smaller data before min we can use asec stack
    for i:=range arr {
        //before first smaller
        for len(stackIndex)>0&&arr[stackIndex[len(stackIndex)-1]]>=arr[i] {
            stackIndex=stackIndex[0:len(stackIndex)-1]//pop
        }
        if len(stackIndex)>0 {
            before[i]=i-stackIndex[len(stackIndex)-1]//first smaller than arr[i]
        }else {
            before[i]=i+1//all not smaller
        }
        stackIndex=append(stackIndex,i)
    }
    fmt.Println(before)
    stackIndex=stackIndex[0:0]//clear
    for i:=len(arr)-1;i>=0;i-- {
        //atfer first smaller
        for len(stackIndex)>0&&arr[stackIndex[len(stackIndex)-1]]>arr[i] {
            //not include =
            stackIndex=stackIndex[0:len(stackIndex)-1]//pop
        }
        if len(stackIndex)>0 {
            after[i]=stackIndex[len(stackIndex)-1]-i//first smaller than arr[i]
        }else {
            after[i]=len(arr)-i
        }
        stackIndex=append(stackIndex,i)
    }
    for i:=0;i<len(arr);i++ {
        res=(res+((after[i]*before[i])%MOD*arr[i])%MOD)%MOD
    }
    return res
}

type Info struct {
    Index int
    Data int
}
