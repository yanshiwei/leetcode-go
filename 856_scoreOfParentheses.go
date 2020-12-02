func scoreOfParentheses(S string) int {
    var res int
    if len(S)<2 {
        return res
    }
    var stack []string
    //mainly is sum*2
    //1finish is ) 
    //.  1.1if head is ( pop and push "1"
    //.  1.2if head is digital, it's (A), sum all digital and * 2 and then pop (
    for i:=range S {
        if S[i] =='(' {
            stack=append(stack,string(S[i]))
        }else {
            //')'
            if stack[len(stack)-1]=="(" {
                // is ()
                stack=stack[0:len(stack)-1]//pop
                stack=append(stack,"1")//push () res is 1
            }else {
                //head is digit
                var sum int
                for len(stack)>0 && stack[len(stack)-1]!="(" {
                    head:=stack[len(stack)-1]
                    stack=stack[0:len(stack)-1]//pop
                    intV,_:=strconv.Atoi(head)
                    sum+=intV
                }
                if len(stack)>0 {
                    stack=stack[0:len(stack)-1]//pop (
                }
                intS:=strconv.Itoa(sum*2)
                stack=append(stack,intS)
            }
        }
    }
    for i:=range stack {
        resTmp,_:=strconv.Atoi(stack[i])
        res+=resTmp
    }
    
    return res
}
