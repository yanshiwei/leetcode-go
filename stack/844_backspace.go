func backspaceCompare(S string, T string) bool {
    var s,t []rune
    for i,oneS:=range S {
        if oneS=='#' {
            if i!=0&&len(s)>0 {
                //not first
                //avoid repeadted # like "##ac"
                s=s[0:len(s)-1]
            }
        }else {
            s=append(s,oneS)
        }
    }
    for i,oneT:=range T {
        if oneT=='#' {
            if i!=0&&len(t)>0 {
                t=t[0:len(t)-1]
            }
        }else {
            t=append(t,oneT)
        }
    }
    if string(s)==string(t) {
        return true
    }else {
        return false
    }
}
