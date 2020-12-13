func decodeString(s string) string {
	var res string
	if len(s)<1 {
		return res
	}
    var st1 []int//num stack
    var st2 []string//str stack
    var sum int
    var str string
    for i:=range s {
        if s[i]-'0'>=0&&s[i]-'0'<=9 {
            //digital
            sum=int(s[i]-'0')+10*sum
        }else {
            if s[i]=='[' {
                st1=append(st1,sum)//push num
                st2=append(st2,str)//push str
                sum=0
                str=""
            }else if s[i]==']' {
                //
                num:=st1[len(st1)-1]//num pop as repeated num
                st1=st1[0:len(st1)-1]
                headStr:=st2[len(st2)-1]//str pop as head string
                st2=st2[0:len(st2)-1]

                for j:=0;j<num;j++ {
                    headStr+=str
                }
                str=headStr
            }else {
                //letter
                str+=string(s[i])
            }
        }
    }
    res=str
    return res
}
