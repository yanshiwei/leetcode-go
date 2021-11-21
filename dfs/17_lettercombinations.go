var res []string
func letterCombinations(digits string) []string {
    var freMap=make([]string,10)
    freMap[0]=" "
    freMap[1]=""
    freMap[2]="abc"
    freMap[3]="def"
    freMap[4]="ghi"
    freMap[5]="jkl"
    freMap[6]="mno"
    freMap[7]="pqrs"
    freMap[8]="tuv"
    freMap[9]="wxyz"
    res=nil
    if len(digits)<1{
        return res
    }
    dfs(digits,0,"",freMap)
    return res
}
func dfs(digits string, idx int, str string, freMap[]string){
    if idx==len(digits){
        // finish
        res=append(res, str)
        return
    }
    c:=digits[idx]
    letters:=freMap[int(c-'0')]
    for i:=0;i<len(letters);i++{
        dfs(digits,idx+1,str+string(letters[i]),freMap)
    }
}
