func average(salary []int) float64 {
    var maxV,sum int
    var minV=INTMAX
    for i:=range salary{
        sum+=salary[i]
        if maxV<salary[i]{
            maxV=salary[i]
        }
        if minV>salary[i]{
            minV=salary[i]
        }
    }
    //fmt.Println(maxV,minV)
    return float64(sum-minV-maxV)/float64(len(salary)-2)
}

const INTMAX=int(^uint(0)>>1)
