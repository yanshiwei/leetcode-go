func buildArray(target []int, n int) []string {
    var res []string
    var length=len(target)
    var i,realIndex int
    for realIndex<length {
        t:=target[realIndex]
        if t==(i+1) {
            res=append(res,"Push")
            realIndex=realIndex+1
            i=i+1
        }else {
            res=append(res,"Push")
            res=append(res,"Pop")
            i=i+1
        }
    }
    return res
}
