func minOperations(boxes string) []int {
    var res=make([]int,len(boxes))
    var abs=func(x int)int {
        if x<0 {
            return -1*x
        }
        return x
    }
    for i:=0;i<len(boxes);i++{
        for j:=0;j<len(boxes);j++{
            if j!=i{
                if boxes[j]=='1'{
                    res[i]+=abs(j-i)
                }
            }
        }
    }
    return res
}
