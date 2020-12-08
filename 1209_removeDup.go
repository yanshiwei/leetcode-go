type Info struct {
    Char byte
    Cnt int
}
func removeDuplicates(s string, k int) string {
    if k<2 {
        return s
    }
    var res string
    var stack[]Info
    for i:=range s {
        if len(stack)>0&&stack[len(stack)-1].Char==s[i] {
            stack[len(stack)-1].Cnt++
            if stack[len(stack)-1].Cnt==k {
                stack=stack[0:len(stack)-1]//pop
            }
        }else {
            info:=Info{Char:s[i],Cnt:1}
            stack=append(stack,info)
        }
    }
    for i:=range stack {
        for j:=0;j<stack[i].Cnt;j++ {
            res+=string(stack[i].Char)
        }
    }
    return res
}

