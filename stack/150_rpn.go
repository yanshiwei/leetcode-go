func evalRPN(tokens []string) int {
    var res int
    if len(tokens)<1 {
        return res
    }
    var stack []string
    for _,one:=range tokens {
        var ress int
        switch one {
            case "+","-","*","/":
                second,_:=strconv.Atoi(stack[len(stack)-1])
                first,_:=strconv.Atoi(stack[len(stack)-2])
                stack=stack[0:len(stack)-2]//pop two
                if one == "+" {
                    ress=first+second
                }else if one == "-" {
                    ress=first-second
                }else if one == "*" {
                    ress=first*second
                }else {
                    ress=first/second
                }
                stack=append(stack,strconv.Itoa(ress))
            default:
                stack=append(stack,one)//push data
        }
    }
    res,_=strconv.Atoi(stack[0])
    return res
}
