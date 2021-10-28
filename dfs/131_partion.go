var res[][]string
var path []string//放已经回文的子串
func partition(s string) [][]string {
    res=nil
    path=nil
    dfs(s,0)
    return res
}

func dfs(str string, startIdx int){
    if(startIdx>=len(str)){
        // 如果起始位置已经大于s的大小，说明已经找到了一组分割方案了
        // 注意只保留path里值。path内容时刻变化
        tmp:=append([]string(nil),path...)
        res=append(res,tmp)
        return
    }
    for i:=startIdx;i<len(str);i++{
        // 获取[startIndex,i]在s中的子串
        if(isPalindrome(str[startIdx:i+1])){
            path=append(path,str[startIdx:i+1])
        }else{
            continue
        }
        dfs(str,i+1)//寻找i+1为起始位置的子串
        path=path[0:len(path)-1]// 回溯过程，弹出本次已经填在的子串
    }
}

func isPalindrome(str string)bool{
    for i,j:=0,len(str)-1;i<j;i,j=i+1,j-1{
        if str[i]!=str[j]{
            return false
        }
    }
    return true
}
