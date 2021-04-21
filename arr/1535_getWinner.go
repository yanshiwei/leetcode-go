func getWinner(arr []int, k int) int {
    //第一遍提交，发现超时，原因是题中k值远大于数组长度，其实当赢得轮数大于数组长度时即可返回；
    /*
    var fre=make(map[int]int)
    var res int
    if k>=len(arr){
        k=len(arr)
    }
    for {
        big,small:=getBigger(arr[0],arr[1])
        copy(arr[0:len(arr)-1],arr[1:])
        arr[0]=big
        fre[big]++
        if fre[big]>=k{
            res= big
            break
        }
        arr[len(arr)-1]=small
    }
    return res
    */
    //无论k有多大，胜利者都在arr数组中，通过一次遍历，如果在碰到数组最大值之前还没有满足条件的胜利者，此时数组最大值来到数组第一位，那么胜利者就必然是数组最大值
    if k>=len(arr){
        k=len(arr)
    }
    var winnerNow=getBigger(arr[0],arr[1])
    var cnt int=1
    for i:=2;i<len(arr);i++{
        if cnt>=k{
            return winnerNow
        }
        if winnerNow<arr[i]{
            winnerNow=arr[i]
            cnt=1
        }else{
            cnt++
        }
    }
    //k超过数组长，最大值
    return winnerNow
}

func getBigger(x,y int)int {
    if x<y {
        return y
    }
    return x
}
