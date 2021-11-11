func isAnagram(s string, t string) bool {
    if len(s)!=len(t){
        return false
    }
    var fre=make(map[byte]*Info)
    for i:=range s{
        if _,ok:=fre[s[i]];ok{
            fre[s[i]].X+=1
        }else{
            fre[s[i]]=&Info{X:1,Y:0}
        }
    }
    for i:=range t{
        if _,ok:=fre[t[i]];ok{
            // exist
            fre[t[i]].Y++
            if fre[t[i]].Y>fre[t[i]].X{
                return false
            }
        }else{
            return false
        }
    }
    return true
}
type Info struct{
    X int
    Y int
}
