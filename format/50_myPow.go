func myPow(x float64, n int) float64 {
    if n<0{
        return 1.0/myPow(x,-1*n)
    }
    if n==0{
        return 1
    }
    var tmp float64=myPow(x,n>>1)
    if n%2==0{
        //偶数
        return tmp*tmp
    }else{
        //奇数
        return tmp*tmp*x
    }
    return 0.0
}
