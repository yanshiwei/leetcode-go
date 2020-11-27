func calculate(s string) int {
	var num []int
	var op []byte
	//current op is smaller or same the pop element that pop to calculate
	for i:=0;i<len(s);i++ {
		if s[i] == ' ' {
			continue
		}
		if int(s[i])-int('0')>=0 && int(s[i])-int('0')<=9 {
			//number
			var tmp int
			var j int
			for j=i;j<len(s);j++ {
				if int(s[j])-int('0')>=0 && int(s[j])-int('0')<=9  {
					tmp=int(s[j])-int('0')+10*tmp
				}else {
					break
				}
			}
			num=append(num,tmp)
			i=j-1//i++
		}else {
			//op str
			if len(op)<1 {
				//empty op arr must be push
				op=append(op,s[i])
			}else {
				//if cur is bigger than head, push
				var head=op[len(op)-1]
				if (head=='+'||head=='-')&&(s[i]=='*'||s[i]=='/') {
					op=append(op,s[i])
				}else {
					//cur is smaller or same than head, pop to cal
					var num1=num[len(num)-1]
					var num2=num[len(num)-2]
					num=num[0:len(num)-2]//pop two
					var op1=op[len(op)-1]
					op=op[0:len(op)-1]//pop op
					var res=cal(num1,num2,op1)
					num=append(num,res)//push res
					//perheps first smaller and after pop become same
					//because same must pop so it must be + -
					if len(op)>0&&(s[i]=='+'||s[i]=='-') {
						num1=num[len(num)-1]
						num2=num[len(num)-2]
						num=num[0:len(num)-2]//pop two
						op1=op[len(op)-1]
						op=op[0:len(op)-1]//pop op
						res=cal(num1,num2,op1)
						num=append(num,res)//push res
					}
					op=append(op,s[i])
				}
			}
		}
	}
	for len(num)>0&&len(op)>0 {
		var num1=num[len(num)-1]
		var num2=num[len(num)-2]
		num=num[0:len(num)-2]//pop two
		var op1=op[len(op)-1]
		op=op[0:len(op)-1]//pop op
		var res=cal(num1,num2,op1)
		num=append(num,res)//push res
	}
	var res int
	if len(num)>0 {
		res=num[len(num)-1]
	}else {
		res=int(s[0]-'0')
	}
	return res
}

func cal(num1,num2 int,op byte)int {
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
