func reverseParentheses(s string) string {
    var res string
    var stack []rune
    for i:=range s{
        if s[i]==')' {
            var tmpArr []rune
            for len(stack)>0&&stack[len(stack)-1]!='(' {
                tmpArr=append(tmpArr,rune(stack[len(stack)-1]))
                stack=stack[0:len(stack)-1]//pop
            }
            if len(stack)>0&&stack[len(stack)-1]=='(' {
                stack=stack[0:len(stack)-1]//pop (
            }
            stack=append(stack,tmpArr...)
        }else {
            stack=append(stack,rune(s[i]))
        }
    }
    res=string(stack)
    return res
}
