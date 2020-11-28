func removeDuplicateLetters(s string) string {
    var strCnt =make(map[rune]int)//26 letter cnt
    var strExist=make(map[rune]bool)
    var res []rune
    for _,one:=range s {
        strCnt[rune(one)]+=1
    }
    //first filter repeated
    //second re-sort by dict sort
    /*
    1首先过滤已经重复的；
    2当栈顶元素大于当前元素且当前元素后面还会储蓄该栈顶元素时，
    将栈顶元素pop一直到条件不成立，才将当前元素入栈。
    3建立一个元素是否存在的判断，在pop栈顶元素时，对那些后面还存在但是已经被第一步过滤的重新加入处理
    */
    for _,one:=range s {
        strCnt[rune(one)]-=1//read one and minus one
        if v,ok:=strExist[rune(one)];ok {
            if v==true {
                continue
            }
            //filter exist already
        }
        if len(res)<1 {
            res=append(res,one)//push if empty
            strExist[one]=true
        }else {
            head:=res[len(res)-1]//head
            if head>=one&&strCnt[head]>0 {
                for head>=one&&strCnt[head]>0 {
                    res=res[0:len(res)-1]//pop
                    strExist[head]=false//so that next can handle
                    if len(res)<1 {
                        break
                    }
                    head=res[len(res)-1]//head
                }
                res=append(res,one)
                strExist[one]=true
            }else {
                res=append(res,one)
                strExist[one]=true
            }
        }
    }
    return string(res)
}
