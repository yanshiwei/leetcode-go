func lengthOfLastWord(s string) int {
    var res int
    var tmp int
    for i:=range s{
        if s[i]==' '{
            if tmp!=0{
                res=tmp
            }
            tmp=0
        }else{
            tmp++
        }
    }
    if tmp!=0{
        res=tmp
    }
    return res
}
