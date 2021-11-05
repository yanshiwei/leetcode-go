func reverseWords(s string) string {
    if len(s)<1{
        return ""
    }
    // delete more space
    news:=removeExtraSpaces(s)
    // reverse the whole str
    reverseStr:=reverseS(news)
    // start handle the word reverse one bye one
    var start int
    var res bytes.Buffer

    for i:=0;i<len(reverseStr);i++{
        if reverseStr[i]==' '{
            oneS:=reverseS(reverseStr[start:i])
            res.WriteString(oneS)
            res.WriteString(" ")
            start=i+1
        }
        if i==len(reverseStr)-1{
            // last word
            oneS:=reverseS(reverseStr[start:])
            res.WriteString(oneS)
        }
    }
    return res.String()
}

func reverseS(str string) string{
    if len(str)<2{
        return str
    }
    var res bytes.Buffer
    for i:=len(str)-1;i>=0;i--{
        res.WriteByte(str[i])
    }
    return res.String()
}

func removeExtraSpaces(str string)string{
    if len(str)<1{
        return str
    }
    //首先删除连续的空格
    var res bytes.Buffer
    for i:=0;i<len(str)-1;i++{
        if str[i]==str[i+1]&&str[i]==' '{
            continue
        }else{
            res.WriteByte(str[i])
        }
    }
    if str[len(str)-1]!=' '{
        res.WriteByte(str[len(str)-1])
    }
    finalres:=res.String()
    //删除字符串末尾的空格
    if len(finalres)>0&&finalres[len(finalres)-1]==' '{
        finalres=finalres[:len(finalres)-1]
    }
    //删除字符串开头的空格
     if len(finalres)>0&&finalres[0]==' '{
        finalres=finalres[1:len(finalres)]
    }   
    return finalres
}
