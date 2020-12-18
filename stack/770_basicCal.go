
func basicCalculatorIV(expression string, evalvars []string, evalints []int) []string {
    index=0 //must have else commit alwalys fail however exe is right
    fmt.Println(expression)
	var res []string
	str = expression
	gEvalvars = evalvars
	gEvalints = evalints
    fmt.Println("before",token,index)
	token = getToken()
     fmt.Println("after",token,index)
	res = exp()
	return res
}

var token string
var str string
var index int
var gEvalvars []string
var gEvalints []int

func getToken() string {
	//找出下一个符号 包括+ - * （ ）digital or val
	var nextToken string
	if index >= len(str) {
		return nextToken
	}
	for len(str) > 0 && str[index] == ' ' {
		index++
		//去除空格
	}

	//digital
	if str[index] >= '0' && str[index] <= '9' {
		//get all digital
		for index < len(str) && (str[index] >= '0' && str[index] <= '9') {
			nextToken += string(str[index])
			index++
		}
	} else if str[index] >= 'a' && str[index] <= 'z' {
		//var
		for index < len(str) && (str[index] >= 'a' && str[index] <= 'z') {
			nextToken += string(str[index])
			index++
		}
	} else {
		//+ - * ( )
		switch str[index] {
		case '+':
			nextToken += "+"
			index++
		case '-':
			nextToken += "-"
			index++
		case '*':
			nextToken += "*"
			index++
		case '/':
			nextToken += "/"
			index++
		case '(':
			nextToken += "("
			index++
		case ')':
			nextToken += ")"
			index++
		}
	}
	return nextToken
}

func isNum(st string) bool {
	for i := range st {
		if st[i] >= 'a' && st[i] <= 'z' {
			return false
		}
	}
	return true
}

//分离因子的系数和变量2*a  2 a
func split(fac1 string) (string, string) {
	var i int
	var numStr, valStr string
	for i < len(fac1) && fac1[i] != '*' {
		i++
	}
	numStr = fac1[0:i]
	if i < len(fac1) {
		valStr = fac1[i+1:] //2*a gap * so is i+1
	}
	return numStr, valStr
}
func splitSub(oriStr, subStr string) []string {
	var res []string
	var start int
	var index = strings.Index(oriStr, subStr)
	for index != -1 && index < len(oriStr) {
		fmt.Println("splitSub", oriStr, start, index)
		res = append(res, oriStr[start:index])
		start = index + len(subStr)
		indexTmp := strings.Index(oriStr[start:], subStr)
		if indexTmp == -1 {
			break
		}
		index = indexTmp + start
	}
	if start < len(oriStr) {
		res = append(res, oriStr[start:])
	}
	return res
}

//处理两个因子相乘 1*a*b*c * 2*c*d
func multi(fac1, fac2 string) string {
	fmt.Println("multi", fac1, fac2)
	var res string
	var numStr1, valStr1, numStr2, valStr2 string
	var num1, num2, coff int
	var valSlice1, valSlice2 []string
	numStr1, valStr1 = split(fac1)
	numStr2, valStr2 = split(fac2)
	num1, _ = strconv.Atoi(numStr1)
	num2, _ = strconv.Atoi(numStr2)
	coff = num1 * num2
	valSlice1 = splitSub(valStr1, "*")
	valSlice2 = splitSub(valStr2, "*")
	sort.Strings(valSlice1) //字典顺序 and has Float54 Ints
	sort.Strings(valSlice2)
	res += strconv.Itoa(coff)
	var i, j int
	for i < len(valSlice1) && j < len(valSlice2) {
		if valSlice1[i] < valSlice2[j] {
			res += "*"
			res += valSlice1[i]
			i++
		} else {
			res += "*"
			res += valSlice2[j]
			j++
		}
	}
	for i < len(valSlice1) {
		res += "*"
		res += valSlice1[i]
		i++
	}
	for j < len(valSlice2) {
		res += "*"
		res += valSlice2[j]
		j++
	}
	return res
}

type NewStrSlice []string

//两个因子比较 1*a*b*c * 2*c*d
func (s NewStrSlice) Len() int      { return len(s) }
func (s NewStrSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s NewStrSlice) Less(i, j int) bool {
	var s1 = s[i]
	var s2 = s[j]
	var num1, num2 int
	for i := 0; i < len(s1); i++ {
		if s1[i] == '*' {
			num1++
		}
	}
	for i := 0; i < len(s2); i++ {
		if s2[i] == '*' {
			num2++
		}
	}
	if num1 != num2 {
		return num1 > num2 //因子度大的优先
	} else {
		var varStr1, varStr2 string
		var i int
		for i < len(s1) && s1[i] != '*' {
			i++
		}
		if i < len(s1) {
			varStr1 = s1[i+1:]
		}
		i = 0
		for i < len(s2) && s2[i] != '*' {
			i++
		}
		if i < len(s2) {
			varStr2 = s2[i+1:]
		}
		return varStr1 < varStr2 //字典顺序优先
	}
}

//合并同类项
func mergeFactor(inputRaw []string) []string {
	var res []string
	if len(inputRaw) < 1 {
		return res
	}
	input := NewStrSlice(inputRaw)
	sort.Sort(input)
	var num int
	var numStr, valStr string
	for i := 0; i < len(input); i++ {
		if i == 0 {
			numStr, valStr = split(input[i])
			num, _ = strconv.Atoi(numStr)
		} else {
			var numStrTmp, valStrTmp string
			numStrTmp, valStrTmp = split(input[i])
			if valStr == valStrTmp {
				//相同的字母
				numTmp, _ := strconv.Atoi(numStrTmp)
				num += numTmp
			} else {
				if num != 0 {
					tp := strconv.Itoa(num)
					if valStr != "" {
						tp += "*" + valStr
					}
					res = append(res, tp)
				}
				numStr = numStrTmp
				valStr = valStrTmp
				num, _ = strconv.Atoi(numStr)
			}
		}
	}
	if num != 0 {
		//插入系数非0的项
		tp := strconv.Itoa(num)
		if valStr != "" {
			tp += "*" + valStr
		}
		res = append(res, tp)
	}
	return res
}

//
func factor() []string {
	var res []string
	if token == "(" {
		token = getToken()
		res = exp()
		token = getToken() //跳过右括号
	} else {
		if isNum(token) {
			res = append(res, token)
		} else {
			for i := 0; i < len(gEvalvars); i++ {
				if token == gEvalvars[i] {
					res = append(res, strconv.Itoa(gEvalints[i]))
				}
			}
			if len(res) < 1 {
				var t string
				t += "1*"
				t += token
				res = append(res, t)
			}
		}
		token = getToken() //跳过num
	}
	return res
}
func additive() []string {
	var res = factor()
	for token == "*" {
		fmt.Println("********")
		token = getToken() //跳过‘*’
		var tmpList []string
		factorList := factor()
		fmt.Println("********", res, factorList)
		for i := 0; i < len(res); i++ {
			for j := 0; j < len(factorList); j++ {
				tmpList = append(tmpList, multi(res[i], factorList[j]))
			}
		}
		res = tmpList
	}
	fmt.Println("additive", res)
	return res
}
func exp() []string {
	var res = additive()
	var tmpList []string
	for token == "+" || token == "-" {
		op := token
		token = getToken() //跳过‘+’ ‘-’
		tmpList = additive()
		if op == "+" {
			for i := 0; i < len(tmpList); i++ {
				res = append(res, tmpList[i])
			}
		} else {
			for i := 0; i < len(tmpList); i++ {
				res = append(res, multi(tmpList[i], "-1"))
			}
		}
	}
	//合并同类项
	var finalRes = mergeFactor(res)
	return finalRes
}


