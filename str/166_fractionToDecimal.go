func fractionToDecimal(numerator int, denominator int) string {
    var n int64=int64(numerator)
    var d int64=int64(denominator)
    var res string
    if n*d<0{
        // 判断正负
        res+="-"
    }
    // 整数部分
    var a int64=abs(n/d)
    res+=strconv.FormatInt(a,10)
    // 分子分母全部变成正号
    if n<0{
        n*=-1
    }
    if d<0{
        d*=-1
    }
    // 取余
    n=n%d
    if n==0{
        return res
    }
    // 添加小数点
    res+="."
    // 判断是否循环
    var index int 
    var s string
    var mp=make(map[int64]int)
    for n>0{
        if _,ok:=mp[n];ok{
            //已经存在，说明有循环
            break
        }
        mp[n]=index
        index++
        n=n*10
        s+=strconv.FormatInt(n/d,10)
        n=n%d
    }
    if n!=0{
        // 发生循环
        idx:=mp[n]//循环开始处
        res+=s[0:idx]+"("+s[idx:]+")"
    }else{
        res+=s
    }
    return res
}

func abs(a int64)int64{
    if a<0{
        return -1*a
    }
    return a
}
