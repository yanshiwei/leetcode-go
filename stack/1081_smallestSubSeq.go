func smallestSubsequence(s string) string {
    var strCnt=make(map[rune]int)
    var strExist=make(map[rune]bool)
    var stack []rune
    for i:=range s {
        strCnt[rune(s[i])]++
    }
    for i:=range s{
        strCnt[rune(s[i])]--
        if v,ok:=strExist[rune(s[i])];ok {
            if v==true {
                continue
            }
        }
        strExist[rune(s[i])]=true
        if len(stack)<1 {
            stack=append(stack,rune(s[i]))
        }else {
            var head=stack[len(stack)-1]
            if len(stack)>0&&head>=rune(s[i]){
                for len(stack)>0&&head>=rune(s[i]) {
                    if strCnt[head]>0 {
                        strExist[head]=false
                        stack=stack[0:len(stack)-1]//pop
                    }else {
                        break
                    }
                    if len(stack)>0 {
                        head=stack[len(stack)-1]
                    }
                    
                }
            }
            stack=append(stack,rune(s[i]))
        }
    }
    return string(stack)
}
