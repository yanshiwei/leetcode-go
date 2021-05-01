func largestAltitude(gain []int) int {
    var maxV int//等于第一个点0的海拔，即为0
    var last int//等于第一个点0的海拔，即为0
    for i:=range gain{
        cur:=gain[i]+last
        last=cur
        if maxV<cur{
            maxV=cur
        }
    }
    return maxV
}
