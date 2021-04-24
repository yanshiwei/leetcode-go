func sumOddLengthSubarrays(arr []int) int {
    var res int
    for i:=0;i<len(arr);i++{
        //i 用来枚举每个连续子数组的起点
        for sz:=1;i+sz-1<len(arr);sz+=2{
            //sz 为连续子数组的长度
            //accumulate用来计算起点是 i，长度为 sz 的子数组的和
            res+=cal(arr[i:i+sz])
        }
    }
    return res
}

func cal(arr []int)int {
    var res int
    for i:=range arr{
        res+=arr[i]
    }
    return res
}
