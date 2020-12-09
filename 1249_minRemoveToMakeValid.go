type Info struct {
    Index int
    Value byte
}
func minRemoveToMakeValid(s string) string {
    var res string
    if len(s)<1 {
        return res
    }
    var stack[]Info
    for i:=range s{
        if s[i]=='(' {
            info:=Info{Index:i,Value:s[i]}
            stack=append(stack,info)
        }else if s[i]==')' {
            if len(stack)>0 &&stack[len(stack)-1].Value=='(' {
                stack=stack[0:len(stack)-1]//pop
            }else {
                info:=Info{Index:i,Value:s[i]}//invalid
                stack=append(stack,info)
            }
        }
    }
    for i:=range s {
        exist:=findIndex(stack,i)
        if exist==false {
            res+=string(s[i])
        }
    }
    return res
}

func findIndex(arr[]Info,index int)bool {
    for i:=range arr {
        if index==arr[i].Index {
            return true
        }
    }
    return false
}
