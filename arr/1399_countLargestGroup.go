func countLargestGroup(n int) int {
    var res int
    var fre=make(map[int]int)
    var maxFre int
    for i:=1;i<=n;i++{
        num:=getSumByStep(i)
        fre[num]++
        if maxFre<fre[num]{
            maxFre=fre[num]
        }
    }
    for _,v:=range fre{
        if v==maxFre{
            res++
        }
    }
    return res
}

func getSumByStep(n int)int {
    var sum int
    for n>0{
        sum+=n%10
        n/=10
    }
    return sum
}
