func entityParser(text string) string {
    var res string
    var stack []byte
    var length=len(text)
    for i:=0;i<length; {
        if text[i]!='&' {
            stack=append(stack,text[i])
            i+=1
        }else {
            if len(text[i+1:])>=5&&text[i+1:i+6]=="quot;" {
                stack=append(stack,'"')
                i+=6
            }else if len(text[i+1:])>=5&&text[i+1:i+6]=="apos;" {
                stack=append(stack,'\'')
                i+=6
            }else if len(text[i+1:])>=4&&text[i+1:i+5]=="amp;" {
                stack=append(stack,'&')
                i+=5
            }else if len(text[i+1:])>=3&&text[i+1:i+4]=="gt;" {
                stack=append(stack,'>')
                i+=4
            }else if len(text[i+1:])>=3&&text[i+1:i+4]=="lt;" {
                stack=append(stack,'<')
                i+=4
            }else if len(text[i+1:])>=6&&text[i+1:i+7]=="frasl;" {
                stack=append(stack,'/')
                i+=7
            }else {
                stack=append(stack,text[i])
                i+=1
            }
        }
    }
    res=string(stack)
    return res
}
