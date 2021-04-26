func countPairs(deliciousness []int) int {
    var fre=make(map[int]int)//数字-出现的次数
    var mod=1000000007
    var res int
    for i:=range deliciousness{
        powerOfTwo:=1
        //因为数字deliciousness[i]最大为2^20,故两个数的和最大为2^21=powerOfTwo
        for j:=0;j<=21;j++{
            if powerOfTwo>=deliciousness[i]{
                if v,ok:=fre[powerOfTwo-deliciousness[i]];ok{
                    res+=v
                    res%=mod
                }
            }
            powerOfTwo*=2
        }
        fre[deliciousness[i]]++
    }
    return res
}
