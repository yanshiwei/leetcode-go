func countBits(n int) []int {
    //奇数：一定比前一个偶数多一个1；偶数：偶数中 1 的个数一定和除以 2 之后的那个数一样多。因为最低位是 0，除以 2 就是右移一位，也就是把那个 0 抹掉而已，所以 1 的个数是不变的
    var res =make([]int,n+1)
    res[0]=0
    for i:=1;i<=n;i++{
        if i%2==1{
            //奇数
            res[i]=res[i-1]+1
        }else{
            res[i]=res[i/2]
        }
    }
    return res
}
