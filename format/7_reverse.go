func reverse(x int) int {
    var res =0
    for x!=0{
        if res<math.MinInt32/10||res>math.MaxInt32/10{
            return 0
        }
        dig:=x%10
        res=res*10+dig
        x/=10
    }
    return res
}
