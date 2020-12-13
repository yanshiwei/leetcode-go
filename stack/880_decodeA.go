func decodeAtIndex(S string, K int) string {
    var res string
    if len(S) <1 {
        return res
    }
    /*
    limit time wrong !!!
    var stack []string
    for i:=0;i<len(S);i++ {
        if S[i]-'0'>=0&&S[i]-'0'<=9 {
            //digital
            var sum=int(S[i]-'0')
            var resTmp=stack
            var k int
            for k=0;k<sum-1;k++ {
                //d-1
                stack=append(stack,resTmp...)
                if len(stack)>=K {
                    break
                }
            }
        }else {
            //letter
            stack=append(stack,string(S[i]))
            if len(stack)>=K {
                break
            }
        }
    }
    fmt.Println(stack)
    res=stack[K-1]
    */
    var curLen int
    var i int
    for i=range S {
        if S[i]-'0'>=0&&S[i]-'0'<=9 {
            curLen*=int(S[i]-'0')
        }else {
            curLen++
        }
    }
    for i>=0 {
        //from tail
        K=K%curLen//len长的字符串扩大d倍后 K处与K%len的值其实是一样的
        if S[i]-'0'>=0&&S[i]-'0'<=9 {
            //如果是数字 则扩大倍数
            curLen=curLen/int(S[i]-'0')
        }else {
            if K==0 {
                return string(S[i])
            }
            curLen--
        }
        i--
    }
    return res
}
