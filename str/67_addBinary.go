func addBinary(a string, b string) string {
    var res string
    var maxLen=maxInt(len(a),len(b))
    var add int
    for i:=0;i<maxLen;i++{
        if i<len(a){
            //reverse
            if a[len(a)-1-i]=='0'{
                add+=0
            }else{
                add+=1
            }
        }
        if i<len(b){
            //reverse
            if b[len(b)-1-i]=='0'{
                add+=0
            }else{
                add+=1
            }
        }
        if add%2==1{
            res+=string('1')
        }else{
            res+=string('0')
        }
        add/=2
    }
    if add>0{
        res+=string('1')
    }
    //reverse
    res=reverse(res)
    return res
}
func maxInt(x,y int)int {
    if x<y{
        return y
    }
    return x
}
func reverse(str string) string {
    var res string
    for i:=0;i<len(str);i++{
        res+=fmt.Sprintf("%c",str[len(str)-i-1])
    }
    return res
}
