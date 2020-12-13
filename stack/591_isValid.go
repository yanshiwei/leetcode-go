func isValid(code string) bool {
    //实际上只有三种类型需要解析：cdata，ending tag和starting tag
    //当遇到一个cdata的起始字符串"<![CDATA["的时候，我们就忽略中间的所有字符，直到找到cdata的特定结尾字符串“]]>”且cdata内容可以是空；
    //档遇到一个ending tag的起始字符串"</"的时候，我们首先解析tag name并判断其合法性，然后查看栈顶元素是否匹配且结束标签不会出现在TAG_CONTENT；
    //当遇到一个starting tag的起始字符串"<"的时候，我们还是首先解析tag name并判断其合法性，如果合法，就将其放入栈中
    if len(code)<5 {
        return false //<></>
    }
    var stackTag []string
    for i:=0;i<len(code); {
        var j int
        if i>0&&len(stackTag)<1 {
            return false//startTag must first
        }
        //start parse startTag
        if len(code[i:])>=9&&code[i:i+9]=="<![CDATA[" {
            //start parse cdata
            j=getSubStrIndex(code[i+9:],"]]>",i+9)//返回]]>第一次出现的位置
            if j<0 {
                return false//根据第7条cdata格式，内容可以是空
            }
            i=j+3
        }else if len(code[i:])>=2&&code[i:i+2]=="</" {
            //start parse endTag
            j=getSubStrIndex(code[i+2:],">",i+2)//返回>第一次出现的位置
            if j<0||i+2==j {
                //根据第4条，内容不含结束标签，且根据第6条必须要有 且非空
                return false
            }
            if j-(i+2)>9 {
                return false//根据第3条必须1-9
            }
            tagName:=code[i+2:j]
            if len(stackTag)<1 {
                return false //根据第5条必须要有startTAG
            }
            if tagName!=stackTag[len(stackTag)-1] {
                return false//根据第2条必须要有一样的TAGNAME
            }
            stackTag=stackTag[0:len(stackTag)-1]//pop
            i=j+1
        }else if len(code[i:])>=1&&code[i:i+1]=="<"{
            //start parse start_tag
            j=getSubStrIndex(code[i+1:],">",i+1)//返回>第一次出现的位置
            if j<0||i+1==j {
                //根据第6条必须要有 且非空
                return false
            }
            if j-(i+1)>9 {
                return false//根据第3条必须1-9
            }
            for k:=i+1;k<j;k++ {
                if code[k]<'A'||code[k]>'Z' {
                    return false //根据第3条必须大写
                }
            }
            tagName:=code[i+1:j]
            stackTag=append(stackTag,tagName)
            i=j+1
        }else {
            i++
        }
    }
    if len(stackTag)>0 {
        //有剩余则不满足8大要求
        return false
    }
    return true
}

func getSubStrIndex(str,subStr string, startIndex int)int {
    if len(str)<=len(subStr) {
        return -1
    }
    j:=strings.Index(str,subStr)//返回>第一次出现的位置
    if j<0 {
        return -1
    }
    return startIndex+j
}
