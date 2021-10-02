func mySqrt(x int) int {
    var res int
    if x<2{
        return x
    }
    for i:=1;i<x;i++{
        if i*i>x{
            break
        }else{
            res=i
        }
    }
    return res
}
