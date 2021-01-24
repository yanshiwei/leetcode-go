func getRow(rowIndex int) []int {
    var res[]int
    if rowIndex<0 {
        return res
    }
    if rowIndex==0{
        res=append(res,1)
        return res
    }
    res=append(res,[]int{1,1}...)
    if rowIndex==1 {
        return res
    }
    for i:=0;i<rowIndex-1;i++ {
        var last=append([]int(nil),res...)
        res=[]int(nil)
        //fmt.Println(last,res)
        res=append(res,1)
        for i:=1;i<len(last);i++ {
            res=append(res,last[i-1]+last[i])
        }
        res=append(res,1)
    }
    return res
}
