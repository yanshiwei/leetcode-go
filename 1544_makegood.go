func makeGood(s string) string {
    if len(s)<2 {
        return s
    }
    var flag=true
    for flag {
        flag=false
        for i:=range s {
            if i>0 {
                if s[i-1]!=s[i] {
                    //big smalle perhapes
                    if strings.ToLower(string(s[i-1]))==strings.ToLower(string(s[i]))                     {
                        s=s[0:i-1]+s[i+1:]
                        flag=true
                        break//one repeat only to handle
                    }
                }
            }
        }
    }
    return s
}
