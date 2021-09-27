func strStr(haystack string, needle string) int {
    var res int
    if len(needle)<1{
        return res
    }
    if len(haystack)<1{
        return -1
    }
    var m=len(needle)
    for i:=0;i<len(haystack);i++{
        if i+m<=len(haystack){
            if needle==haystack[i:i+m]{
                return i
            }
        }else{
            return -1
        }
    }
    return -1
}
