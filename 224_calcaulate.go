    //有数字了就尝试及时计算保证+ - 的先后顺序
    //有右括号了去左括号并及时计算保证括号优先顺序
func calculate(s string) int {
    var res int
    var dataSt []int
    var opSt []byte
    for i:=0;i<len(s);{
        if s[i]==' ' {
            fmt.Println(i)
            i++
            continue
        }else if s[i]=='('||s[i]=='+'||s[i]=='-' {
            opSt=append(opSt,s[i])
            i++
        }else if s[i]==')' {
            opSt=opSt[0:len(opSt)-1]//pop left (
            //尝试计算 例如1+（2-3）
            if len(opSt)>0&&opSt[len(opSt)-1]!='(' {
                var num1=dataSt[len(dataSt)-1]
	            var num2=dataSt[len(dataSt)-2]
	            dataSt=dataSt[0:len(dataSt)-2]//pop two
	            var op=opSt[len(opSt)-1]
	            opSt=opSt[0:len(opSt)-1]//pop op
                var calRes int
	            if op=='+' {
		            calRes= num2+num1
	            }else if op=='-' {
		            calRes= num2-num1
	            }
                dataSt=append(dataSt,calRes)
            }
            i++
        }else {
            var sum int
            var j int
            for j=i;j<len(s);j++ {
                if s[j]-'0'>=0&&s[j]-'0'<=9 {
                    data:=int(s[j]-'0')
                    sum=data+sum*10
                }else {
                    break
                }
            }
            fmt.Println(sum)
            i=j
            dataSt=append(dataSt,sum)
            //尝试计算如2-1+2，得到1后先做减，及时计算保证先后顺序
            if len(opSt)>0&&opSt[len(opSt)-1]!='(' {
                var num1=dataSt[len(dataSt)-1]
	            var num2=dataSt[len(dataSt)-2]
	            dataSt=dataSt[0:len(dataSt)-2]//pop two
	            var op=opSt[len(opSt)-1]
	            opSt=opSt[0:len(opSt)-1]//pop op
                var calRes int
	            if op=='+' {
		            calRes= num2+num1
	            }else if op=='-' {
		            calRes= num2-num1
	            }
                dataSt=append(dataSt,calRes)
            }
        }
    }
    
    if len(dataSt)>0 {
		res=dataSt[len(dataSt)-1]
	}else {
		res=int(s[0]-'0')
	}
    return res
}
func cal(dataSt []int,opSt []byte)int {
    fmt.Println(dataSt)
    var num1=dataSt[len(dataSt)-1]
	var num2=dataSt[len(dataSt)-2]
	dataSt=dataSt[0:len(dataSt)-2]//pop two
	var op=opSt[len(opSt)-1]
	opSt=opSt[0:len(opSt)-1]//pop op
	if op=='+' {
		return num2+num1
	}else if op=='-' {
		return num2-num1
	}else if op =='*' {
		return num2*num1
	}else if op=='/'{
		return num2/num1
	}
	return 0
}
