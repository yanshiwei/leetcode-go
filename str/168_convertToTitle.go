func convertToTitle(columnNumber int) string {
    var res[]byte
    for columnNumber>0{
        var ress=columnNumber%26
        if ress==0{
            //如果余数是0，就像上一位借个1（26）出来，让余数强行等于26
            ress=26
            columnNumber-=26
        }
        res=append(res,byte(64+ress))
        columnNumber/=26
    }
    final:=string(res)
    return reverse(final)
}

func reverse(str string) string {
    var result string
    strLen := len(str)
    for i := 0; i < strLen; i++ {
        result = result + fmt.Sprintf("%c", str[strLen-i-1])
    }
    return result
}
