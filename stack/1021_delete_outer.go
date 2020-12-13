func removeOuterParentheses(S string) string {
    var tmpList []string
    if len(S)<2 {
        return ""
    }
    var st []rune
    var lastEnd int
    for i,s:=range S{
        if s=='(' {
            st=append(st,s)//push
        }else {
            //s is `(`
            st=st[0:len(st)-1]//pop
            if len(st)==0 {
                //one primitive
                tmpList=append(tmpList,S[lastEnd:i+1])
                lastEnd=i+1
            }
        }
    }
    var buf bytes.Buffer//more efficient
    for _,one:=range tmpList {
        newOne:=deleteOuterParenthess(one)
        buf.WriteString(newOne)
    }
    return buf.String()
}

func deleteOuterParenthess(S string) string {
    if len(S)<2 {
        return ""
    }
    return S[1:len(S)-1]
}
