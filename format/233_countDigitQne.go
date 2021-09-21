func countDigitOne(n int) int {
    var i int=1//i表示记录到了第几位
    var sum int
    var befor=0//befor身后的数字
    for n>0{
        tmp:=n%10//记录这个位置是什么数字
        n/=10//n此时是这个位置前面的数字
        if tmp>1{
            //如果这个位置大于1，前面的位置有n+1种取法
            sum+=(n+1)*i
        }else if tmp==1{
            //等于1，前面的位置有n种取法+后面的位置tmp*i种+当前
            sum+=n*i+befor+1
        }else{
            //这个位置是0，则前面位置只有n种取法
            sum+=n*i
        }
        befor+=tmp*i//记录身后的数字
        i*=10//next位
    }
    return sum
}
