func isHappy(n int) bool {
    if n==1{
        return true
    }
    slow,fast:=n,next(n)
    for fast!=1&&slow!=fast{
        slow=next(slow)
        fast=next(next(fast))
    }
    return fast==1
}
func next(n int)int {
    var res int
    for n>0{
        tmp:=n%10
        res+=(tmp*tmp)
        n=n/10
    }
    return res
}
