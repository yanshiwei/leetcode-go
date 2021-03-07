func numPairsDivisibleBy60(time []int) int {
    /*
    成对的时间和为60整除，则成对的时间mod60后余数之和必为60.
    对所有时间求 mod 60 的余数，然后进行归类计数，余数相同的为一类
    除以余数为0的，之和自己同类的其他时间配对
    设一个余数设a，一个余数b，满足a+b=60，a对应数目m，b对应数目n,则配对数m*n/2，（除以2去掉重复）
    */
    var fre=make([]int,60)
    for i:=range time{
        fre[time[i]%60]++
    }
    var res int
    for i:=range fre{
        if fre[i]==0{
            //不存在，则跳过
            continue
        }
        if i==0||i==30{
            //余数0的，与自己同一类的配对
            res+=fre[i]*(fre[i]-1)
        }else{
            res+=fre[i]*fre[60-i]
        }
    }
    //遍历数组进行统计求和会发生同一对的二次计数问题
    res=res>>1//除以2
    return res
}
