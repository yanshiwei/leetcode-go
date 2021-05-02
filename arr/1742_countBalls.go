func countBalls(lowLimit int, highLimit int) int {
    getSum:=func(x int)int{
        var sum int
        for x>0{
            mod:=x%10
            sum+=mod
            x/=10
        }
        return sum
    }
    var fre=make(map[int]int)
    var maxFre int
    for i:=lowLimit;i<=highLimit;i++{
        sum:=getSum(i)
        fre[sum]++
        if maxFre<fre[sum]{
            maxFre=fre[sum]
        }
    }
    return maxFre
}
