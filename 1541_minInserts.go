func minInsertions(s string) int {
    var ans int
    var stack[]byte
    for i:=0;i<len(s);{
        if s[i]=='(' {
            stack=append(stack,s[i])
            i+=1
        }else {
            //right
            if len(stack)==0 {
                //add left
                ans+=1
            }else {
                stack=stack[0:len(stack)-1]//pop
            }
            if i+1==len(s)||s[i+1]!=')' {
                //add right
                i+=1
                ans+=1
            }else {
                i+=2
            }
        }
    }
    ans+=2*len(stack)//stack contains all left
    return ans
}
