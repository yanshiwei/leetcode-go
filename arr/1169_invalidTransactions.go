func invalidTransactions(transactions []string) []string {
	var res []string
	for i:=range transactions{
		oneTrans:=transactions[i]
		oneContent:=strings.Split(oneTrans,",")
		money:=aToi(oneContent[2])
		if money>1000{
			res=append(res,oneTrans)
			continue
		}
		city:=oneContent[3]
		name:=oneContent[0]
		time:=aToi(oneContent[1])
		for j:=range transactions{
			if i==j{
				continue
			}
			anoTrans:=transactions[j]
			anoContent:=strToarr(anoTrans)
			cityAno:=anoContent[3]
			nameAno:=anoContent[0]
			timeAno:=aToi(anoContent[1])
			if name==nameAno&&city!=cityAno&&abs(time-timeAno)<=60{
				res=append(res,oneTrans)//not another
				break
			}
		}
	}
	return res
}

func abs(z int)int {
	if z<0 {
		return -1*z
	}
	return z
}

func aToi(a string)int {
	ai,_:=strconv.Atoi(a)
	return ai
}

func strToarr(a string)[]string{
	return strings.Split(a,",")
}
