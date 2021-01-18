func plusOne(digits []int) []int {
    var first =true
    var add bool
    var res []int
    for i:=len(digits)-1;i>=0;i--{
        one:=digits[i]
        var tmp =one
        if first {
            tmp=one+1
            first=false
        }else if add {
            tmp=one+1
        }
        if tmp>=10{
            add=true
            tmp-=10 
        }else {
            add=false
        }
        digits[i]=tmp
    }
    if add {
        res=append(res,1)
    }
    res=append(res,digits...)
    return res
}
