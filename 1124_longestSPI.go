func longestWPI(hours []int) int {
    var res int
    if len(hours)<1 {
        return res
    }
    //hours: 9  9  6  0  6  6  9
    //ht.  : 1  1 -1 -1 -1 -1  1
    var visited =make(map[int]int)//每个前缀和最早出现的位置
    //最优解sum只可能是1 反证法：如果大于1，说明可以继续向后走就不是最优解了
    var sum,realV int
    for i:=range hours {
        if hours[i]>8 {
            realV=1
        }else {
            realV=-1
        }
        sum+=realV
        if sum>0 {
            res=i+1
        }else {
            if _,ok:=visited[sum-1];ok {
                //(sum)-(sum-1)=1 出现了0，则之前肯定出现过-1，-1经过一系列加减最后是0，这一系列加减总效果=1   故需要计算-1到0这段距离
                res=max(res,i-visited[sum-1])
            }
            if _,ok:=visited[sum];ok==false {
                //记录负数第一次出现的下标
                visited[sum]=i
            }
        }
    }
    return res
}
func max(x,y int)int {
    if x<y{
        return y
    }
    return x
}
