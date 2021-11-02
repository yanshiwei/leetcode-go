func isIsomorphic(s string, t string) bool {
    var lens=len(s)
    if lens<1{
        return false
    }
    var lent=len(t)
    if lent<1{
        return false
    }
    if lens!=lent{
        return false
    }
    var fre=make(map[byte]byte)
    var mirror=make(map[byte]byte)
    for i:=0;i<lens;i++{
        if v,ok:=fre[s[i]];ok{
            if t[i]!=v{
                return false
            }
        }else{
            if _,ok:=mirror[t[i]];ok{
                // already pointed by others
                return false
            }
            fre[s[i]]=t[i]
            mirror[t[i]]=s[i]
        }
    }
    return true
}
