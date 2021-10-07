func hammingWeight(num uint32) int {
    var res int
    var tmp uint32 =1
    for i:=0;i<32;i++{
        if num&tmp>0{
            res++
        }
        tmp=tmp<<1
    }
    return res
}
