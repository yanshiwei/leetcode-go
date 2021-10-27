func reverseBits(num uint32) uint32 {
    var res uint32
    for i:=0;i<32&&num>0;i++{
        res|=(num&1)<<(31-i)//放到高位
        num>>=1//每枚举一位就将 n 右移一位,这样当前 n 的最低位就是我们要枚举的比特位
    }
    return res
}
