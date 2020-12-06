func minAddToMakeValid(S string) int {
    var res int
    if len(S)<1 {
        return res
    }
    var stack []string
    for i:=range S {
        if S[i]=='(' {
            stack=append(stack,string('('))
        }else {
            //)
            if len(stack)==0 {
                res++
            }else {
                stack=stack[0:len(stack)-1]//pop
            }
        }
    }
    if len(stack)>0 {
        res+=len(stack)
    }
    return res
}
