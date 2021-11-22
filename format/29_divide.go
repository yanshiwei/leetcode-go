func divide(dividend int, divisor int) int {
    //https://leetcode-cn.com/problems/divide-two-integers/solution/geekplayers-leetcode-ac-qing-xi-yi-dong-05tn2/
    var m int64=abs(int64(dividend))
    var n int64=abs(int64(divisor))
    if dividend==-2147483648&&divisor==-1{
        return 2147483647
    }
    if dividend==-2147483648&&divisor==1{
        return -2147483648
    }
    var res int64
    var sign int
    if dividend<0&&divisor<0||dividend>0&&divisor>0{
        sign=1
    }else{
        sign=-1
    }
    for m>=n{
        var d int64=n
        var c int64=1
        // 除数n的倍数d, 计数c
        for m>=(d<<1){
            d<<=1// d=n*(2^i), 每次扩大2倍
            c<<=1// c=2^i，与d同步加倍   
        }
        res+=c
        m-=d
    }
   if res < -2147483648||res>2147483647{
        return 2147483647
    }
    if sign==1{
        return int(res)
    }
    return -1*int(res)
}
func abs(a int64)int64{
    if a<0{
        return -1*a
    }
    return a
}
