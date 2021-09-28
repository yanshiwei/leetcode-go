func longestCommonPrefix(strs []string) string {
    var res string
    if len(strs)<1{
        return res
    }
    res=strs[0]
    for i:=1;i<len(strs);i++{
        res=lcs(res,strs[i])
        if len(res)<1{
            break
        }
    }
    return res
}

func lcs(str1,str2 string)string{
    var length=min(len(str1),len(str2))
    var idx int
    for idx<length{
        if str1[idx]!=str2[idx]{
            break
        }
        idx++
    }
    return str1[0:idx]
}
func min(x,y int)int{
    if x<y{
        return x
    }
    return y
}
