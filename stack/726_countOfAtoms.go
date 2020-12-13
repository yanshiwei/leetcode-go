type atomNameSlice []string
func (s atomNameSlice)Len()int {return len(s)}
func (s atomNameSlice)Swap(i,j int){s[i],s[j]=s[j],s[i]}
func (s atomNameSlice)Less(i,j int)bool {
    if s[i]<s[j] {
        return true
    }
    return false
}
//https://www.bilibili.com/video/BV1ex411V79y?from=search&seid=4643607089062262728
func countOfAtoms(formula string) string {
    //一个括号或非括号内容对应一个map 
    var res string
    if len(formula)<1 {
        return res
    }
    var stack[]map[string]int//map key is atom name value is cnt
    var atomName atomNameSlice
    var curMap=make(map[string]int)
    for i:=0;i<len(formula); {
        if formula[i]=='(' {
            stack=append(stack,curMap)
            curMap=make(map[string]int)//clear
            i++
        }else if formula[i]==')' {
            tmpMap:=curMap
            curMap=stack[len(stack)-1]
            stack=stack[0:len(stack)-1]//pop
            //merge history map and current map
            i++
            var oriStart=i
            var factor=1
            for i<len(formula)&&isDigita(formula[i]) {
                i++
            }
            if oriStart<i {
                factor,_=strconv.Atoi(formula[oriStart:i])
            }
            for k,v:=range tmpMap {
                if _,ok:=curMap[k];ok==false {
                    //new
                    curMap[k]=v*factor
                }else {
                    curMap[k]+=v*factor
                }
            }

        }else {
            //atom
            var oriStart=i//首字母是大写
            i++
            for i<len(formula)&&isSmallLetter(formula[i]) {
                i++
            }
            atomName:=formula[oriStart:i]
            
            oriStart=i
            var factor=1
            for i<len(formula)&&isDigita(formula[i]) {
                i++
            }
            if oriStart<i {
                factor,_=strconv.Atoi(formula[oriStart:i])
            }
            fmt.Println("map",atomName,factor)
            if _,ok:=curMap[atomName];ok==false {
                curMap[atomName]=factor
            }else {
                curMap[atomName]+=factor
            }
        }
    }
    //final res is at curMap
    for k,_:=range curMap {
        atomName=append(atomName,k)
    }
    sort.Sort(atomName)
    for _,name:=range atomName{
        res+=name
        if curMap[name]>1 {
            //大于1才展示数字
            res+=strconv.Itoa(curMap[name])
        }
    }
    return res
}
func isBigLetter(a byte)bool {
    if a>='A'&&a<='Z' {
        return true
    }
    return false
}
func isDigita(a byte)bool {
    if a>='0'&&a<='9' {
        return true
    }
    return false
}
func isSmallLetter(a byte)bool {
    if a>='a'&&a<='z' {
        return true
    }
    return false
}
