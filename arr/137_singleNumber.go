func singleNumber(nums []int) int {
    /*
    将每个数想象成32位的二进制，对于每一位的二进制的1和0累加起来必然是3N或者3N+1， 为3N代表目标值the only one在这一位没贡献，3N+1代表目标值the only one在这一位有贡献(=1)，然后将所有有贡献的位|起来就是结果。这样做的好处是如果题目改成K个一样，只需要把代码改成cnt%k，很通用
    */
    var res int
    if len(nums)<1{
        return 0
    }
    var fre=make(map[int]int)
    for _,v:=range nums{
        fre[v]++
    }
    for k,v:=range fre{
        if v==1{
            res=k
            break
        }
    }
    return res
}
