func convert(s string, numRows int) string {
    if numRows==1{
        return s
    }
    //使用min(numRows,len(s)) 个列表来表示 Z 字形图案中的非空行
    var rows=make([]string,min(numRows,len(s)))
    var curRow int
    var gogingDown bool
    for i:=range s{
        rows[curRow]+=string(s[i])
        //get final direct
        if curRow==0||curRow==numRows-1{
            //最开始和最末，肯定转向
            gogingDown=!gogingDown
        }
        if gogingDown{
            curRow+=1
        }else{
            curRow-=1
        }
    }
    var res string
    for i:=range rows{
        res+=rows[i]
    }
    return res
}
func min(x,y int)int {
    if x<y{
        return x
    }
    return y
}
