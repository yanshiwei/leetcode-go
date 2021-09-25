var res []string
func generateParenthesis(n int) []string {
    res=nil
    if n<1{
        return res
    }
    dfs("",n,0,0)
    return res
}
func dfs(paths string,n,left,right int){
    if left>n || right>left{
        //我们生成同一方向的括号只需要n个即可
        //剩下的右括号比左括号少时，是不能匹配
        return
    }
    if len(paths)== n*2{
        res=append(res,paths)
        return
    }//因为括号都是成对出现的
    dfs(paths+"(",n,left+1,right)
    dfs(paths+")",n,left,right+1)
}
