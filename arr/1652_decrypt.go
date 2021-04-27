func decrypt(code []int, k int) []int {
    var res =make([]int,len(code))
    if k==0{
        return res
    }
    var realk=k
    if realk<0{
        k=-1*k
    }
    for i:=0;i<len(code);i++{
        var sum int
        var curIdx=i
        for j:=0;j<k;j++{
            if realk>0{
                curIdx=(curIdx+1)%len(code)
            }else{
                curIdx=(curIdx-1+len(code))%len(code)
            }
            //fmt.Println(curIdx)
            sum+=code[curIdx]
        }
        res[i]=sum
    }
    return res
}
