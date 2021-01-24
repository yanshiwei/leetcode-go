func generate(numRows int) [][]int {
    var res[][]int
    for i:=0;i<numRows;i++{
        var one []int
        if len(res)<1 {
            one=append(one,1)
        }else if len(res)==1{
            one=append(one,1)
            one=append(one,1)
        }else {
            var last=res[i-1]
            one=append(one,1)
            for i:=1;i<len(last);i++ {
                one=append(one,last[i-1]+last[i])
            }
            one=append(one,1)
        }
        res=append(res,one)
    }
    return res
}
