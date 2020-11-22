func removeDuplicates(S string) string {
    if len(S)<2 {
        return S
    }
    var flag=true
    for flag!=false {
        flag=false
        for i:=range S {
            if i>0 {
                if S[i-1]==S[i] {
                    //one only
                    S=S[0:i-1]+S[i+1:]
                    flag=true
                    break
                }
            }
        }
    }
    return S 
}
