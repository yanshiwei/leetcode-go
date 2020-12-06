func isValid(s string) bool {
    if len(s)<3 {
        return false
    }
    if len(s)%3!=0 {
        return false
    }
    if s[0]!='a' {
        return false
    }
    var stack []byte
    for i:=range s {
        if s[i]=='c' {
            //must c end 
            //pop two must b and a
            if len(stack)<1 {
                return false
            }else if stack[len(stack)-1]!='b' {
                return false
            }else {
                stack=stack[0:len(stack)-1]//pop a
                if len(stack)<1 {
                    return false
                }else if stack[len(stack)-1]!='a' {
                    return false
                }
                stack=stack[0:len(stack)-1]//pop b
            }
        }else {
            stack=append(stack,s[i])
        }
    }
    return true
}
