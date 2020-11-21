//stack left push and right pop
func isValid(s string) bool {
    if len(s)<2 {
        return false
    }
    var stack []byte
    for i:=0;i<len(s);i++{
        one:=s[i]
        switch one {
        case '(','[','{':
            stack=append(stack,one)//push
        case ')',']','}':
            //right must has left in stack
            if len(stack)<1 {
                return false
            }
            if one==')' && stack[len(stack)-1]=='(' ||
              one==']' && stack[len(stack)-1]=='[' || 
              one=='}' && stack[len(stack)-1]=='{' {
                  stack=stack[0:len(stack)-1]//pop
              }else {
                  return false
              }
        default:
            //othter commons
            return false
        }
    }
    if len(stack)>0 {
        // res commons
        return false
    }
    return true
}

